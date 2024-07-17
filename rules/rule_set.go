package rules

import utilSet "github.com/dozer111/projectlinter-core/util/set"

type (
	// Set Ряд правил згрупований під одним іменем(composer,werf,phpunit, ...)
	//
	// Правила повертаються саме як RuleTree, тому що на відміну від багатьох інших парсерів(phpCsFixer, rector, golangci)
	// цей не перевіряє окремі ділянки кода. Він перевіряє структуру(конфігурацій/директорій) загалом,
	// тому нормально що є правила більш глобальні, а є більш точкові.
	// Set A list of rules grouped by some thing(composer,phpunit, go.mod, ...)
	//
	// Rules are returned exactly as a RuleTree, because unlike many other linters(phpCsFixer, rector, golangci)
	// this one does not check individual sections of code. It checks the configuration files(or directories)  in complex.
	// It means that this parser knows that there is a primary rules and a secondary(dependant by primary) rules
	// So, it make sense to firstly check only primary rule(s), and if it correct - check the secondary
	// Example:
	//
	// We have a correct Makefile. Let's imagine that new version if it also use new fileName in .helm directory in our project
	// So, projectlinter need to check 2 things
	// 1. Makefile is latest
	// 2. There is exactly file with new name in .helm directory
	//
	// The realization of set make it possible to firstly check "Makefile is latest", and only if its true - do the second
	// check "<fileWithNewName> exists in .helm directory"
	Set interface {
		// ID to ignore via rules + це значення завжди буде коротеньке, тому воно слугує ще й іменем
		ID() string
		// Init actually prepares everything necessary before Run(parse the corresponding configuration files, ...)
		Init() []error
		// Run - actually execute the rules
		Run() *RuleTree
	}

	// Linter is actually a struct(thing) that checks your project
	//
	// In general, this is the same as when you do
	// phpCsFixer fix
	// whether rector run
	// ...
	//
	// You collect a list of rule Set. And run all of them
	Linter interface {
		Sets() ([]Set, error)
	}

	// RuleTree - паттерн composite
	//
	// Структура для зберігання і роботи з правилами
	// Завдяки своїй гнучкості дозволяє створювати не просто прості лінійні набори правил(як наприклад в golangci,phpcsfixer,rector, ...)
	//
	// А доволі складні структури що мають кілька рівнів вкладеності
	//
	// Чому саме так - тому що щоб сконфігурувати projectlinter вам треба писати код(а не просто зробити масив чи yaml файл)
	// Тому, оскільки ми вже все одно маємо докладати зусилля для конфігурації ми маємо отримувати певний виграш
	//
	// Виграш вкладеності заключається в тому що ви виставляєте знання про те як налаштувати свій проект більш правильно
	// а не просто хаотичні зауваження в стилі "зміни те на це"
	// RuleTree - composite pattern
	//
	// Structure for storing and working with rules
	// Thanks to its flexibility, it allows you to create not just simple linear sets of rules (such as in golangci, phpcsfixer, rector, ...)
	//
	// And rather complex structures that have several nesting levels
	//
	// Why exactly so? - because you need to write code to configure your projectlinter cli app (not just make an array or a yaml file)
	// Therefore, since we still have to make efforts for the configuration, we should get some profit
	//
	// The benefit of nesting is that you expose knowledge of how to configure your project more correctly
	// not just random checks like "change this to this"
	RuleTree struct {
		Rules []RuleTreeLeaf
	}

	RuleTreeLeaf struct {
		Rule     Rule
		Children []RuleTreeLeaf
		// Conditions additional conditions that must be met in order for the rule to be ready for verification
		// Use when similar checks cannot (or are not rational) to be done due to nested Children
		// Example
		//
		// Lets imagine that i have 5 rules: rule1, rule2, rule3, rule4, rule5
		//
		// The rules nesting is
		// 	rule1
		//		rule2
		//		rule3
		//	rule4
		//
		// As you see, rule4 is not related to rule1(is not a child of rule1)
		// I want to run rule5 only if rules 1,2,3,4 are done
		// So, without conditions I cannot do it
		// With the conditions I can do something like
		// rule5.Conditions = []func()bool{
		//	rule1.IsPassed,
		//	rule2.IsPassed,
		//	rule3.IsPassed,
		//	rule4.IsPassed,
		//}
		//
		//
		//
		Conditions []func() bool
		// Optional вказує на те чи є правило опційним.
		// Якщо правило опційне - воно та його Children не попадуть в Resolve як "непройдені" в разі невдачі
		// Optional indicates whether the rule is optional.
		// If the rule is optional and its fail - projectlinter would pretend that this rule was newer in list
		//
		// This is cool feature
		// The sense is that sometimes your configuration may have a suitable block of code.
		// So, if the block exists - we need to check it.
		// If block is absent - projectlinter must not mark it as fail. It must just skip it
		//
		// Real example:
		// In my practice I write the integration PHP tests on some services
		// This affect my CI scripts because I need to pass docker socket
		//
		// So, If my project has integration test - I need to check that CI script also check that docker.sock is passed
		// If my project has no integration tests - that's also OK. I just skip this leaf(with all the children) because
		// as u see now - this block is not necessary for me
		Optional bool
	}
)

func NewRuleTree(rules ...RuleTreeLeaf) *RuleTree {
	return &RuleTree{
		rules,
	}
}

func NewLeaf(r Rule, children ...RuleTreeLeaf) RuleTreeLeaf {
	return RuleTreeLeaf{
		Rule:     r,
		Children: children,
	}
}

func NewOptionalLeaf(r Rule, children ...RuleTreeLeaf) RuleTreeLeaf {
	return RuleTreeLeaf{
		Rule:     r,
		Optional: true,
		Children: children,
	}
}

func NewLeafWithConditions(r Rule, conditions []func() bool, children ...RuleTreeLeaf) RuleTreeLeaf {
	if len(conditions) == 0 {
		panic("conditions cannot be empty")
	}

	return RuleTreeLeaf{
		Rule:       r,
		Conditions: conditions,
		Children:   children,
	}
}

// Resolve - ignore - list of rule/set IDs
func (t *RuleTree) Resolve(ignore []string) []Rule {
	ignoreSet := utilSet.NewSet[string](ignore...)
	rules := make([]Rule, 0, len(t.Rules)*2)

	for _, l := range t.Rules {
		rules = t.resolveRules(rules, l, ignoreSet)
	}

	return rules
}

func (t *RuleTree) resolveRules(rules []Rule, leaf RuleTreeLeaf, ignore *utilSet.Set[string]) []Rule {
	conditionsAreDone := true

	for i := 0; i < len(leaf.Conditions); i++ {
		if !leaf.Conditions[i]() {
			conditionsAreDone = false
			break
		}
	}

	if !conditionsAreDone {
		return rules
	}

	rule := leaf.Rule

	if ignore.Has(rule.ID()) {
		return rules
	}

	rule.Validate()

	// if rule is not passed and its optional - does not resolve and return it
	// just pretend that this rule was not add
	if !rule.IsPassed() && leaf.Optional {
		return rules
	}

	rules = append(rules, rule)
	if rule.IsPassed() {
		for _, c := range leaf.Children {
			rules = t.resolveRules(rules, c, ignore)
		}
	}

	return rules
}
