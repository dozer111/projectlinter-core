package dependency

import (
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/golang/gomod/config"
	"github.com/dozer111/projectlinter-core/rules/golang/gomod/parser"
	"github.com/dozer111/projectlinter-core/util/path_provider"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
)

// GolangDependencySet reusable already created RuleSet
//
// # It uses already generated bump configs to check
//
// As for me - the code of dependency module(dependencySet) would always the same
// So no need to write some custom - use this instead
type GolangDependencySet struct {
	pathProvider *path_provider.PathProvider

	substituteRules        []substitute.Library
	substituteDependencies []substitute.Dependency

	bumpRules        []bump.Library
	bumpDependencies []bump.Dependency

	id string
}

var _ rules.Set = (*GolangDependencySet)(nil)

func NewGolangDependencySet(
	pathProvider *path_provider.PathProvider,
	substituteRules []substitute.Library,
	bumpRules []bump.Library,
) *GolangDependencySet {
	return &GolangDependencySet{
		pathProvider:    pathProvider,
		substituteRules: substituteRules,
		bumpRules:       bumpRules,
		id:              "dependency_golang",
	}
}

func (s *GolangDependencySet) ID() string {
	return s.id
}

func (s *GolangDependencySet) SetID(name string) {
	s.id = name
}

func (s *GolangDependencySet) Init() []error {
	gomodParser := parser.NewParser(s.pathProvider.PathToCaller())
	gomodConfig, err := gomodParser.Parse()
	if err != nil {
		return []error{err}
	}

	cfg := gomodConfig

	s.bumpDependencies = goDependenciesToBumpDependencies(cfg.Dependencies)
	s.substituteDependencies = goDependenciesToSubstituteDependencies(cfg.Dependencies)

	return nil
}

func (s *GolangDependencySet) Run() *rules.RuleTree {
	return rules.NewRuleTree(
		leaf(
			substitute.NewSubstituteGOLibraryRule(
				s.ID(),
				s.substituteRules,
				s.substituteDependencies,
			),
		),
		leaf(
			bump.NewBumpGOLibraryRule(
				s.ID(),
				s.bumpRules,
				s.bumpDependencies,
			),
		),
	)
}

// leaf - this is commonly used practise for sets in projectlinter
//
// # The main reason is - readability
//
// As for me - it is easier to read the code like
// NewRuleTree
//
//	leaf
//	leaf
//		leaf
//		leaf
//	leaf
//	leaf
//		leaf
//	leaf
//
// instead of
//
// NewRuleTree
//
//	rules.NewLeaf
//	rules.NewLeaf
//		rules.NewLeaf
//		rules.NewLeaf
//	rules.NewLeaf
//
// ...
func leaf(r rules.Rule, children ...rules.RuleTreeLeaf) rules.RuleTreeLeaf {
	return rules.NewLeaf(r, children...)
}

func goDependenciesToBumpDependencies(dependencies *config.GomodDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		for _, concreteDependency := range v {
			result = append(result, concreteDependency)
		}
	}

	return result
}

func goDependenciesToSubstituteDependencies(dependencies *config.GomodDependencies) []substitute.Dependency {
	result := make([]substitute.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		for _, concreteDependency := range v {
			result = append(result, concreteDependency)
		}
	}

	return result
}
