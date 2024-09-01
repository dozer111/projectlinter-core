package dependency

import (
	"github.com/dozer111/projectlinter-core/rules"
	composerConfig "github.com/dozer111/projectlinter-core/rules/php/composer/config"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/parser"
	"github.com/dozer111/projectlinter-core/util/path_provider"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
)

// PHPDependencySet reusable already created RuleSet
//
// # It uses already generated bump configs to check
//
// As for me - the code of dependency module(dependencySet) would always the same
// So no need to write some custom - use this instead
type PHPDependencySet struct {
	pathProvider *path_provider.PathProvider

	substituteRules        []substitute.Library
	substituteDependencies []substitute.Dependency

	bumpRules        []bump.Library
	bumpDependencies []bump.Dependency

	id string
}

var _ rules.Set = (*PHPDependencySet)(nil)

func NewPHPDependencySet(
	pathProvider *path_provider.PathProvider,
	substituteRules []substitute.Library,
	bumpRules []bump.Library,
) *PHPDependencySet {
	return &PHPDependencySet{
		pathProvider:    pathProvider,
		substituteRules: substituteRules,
		bumpRules:       bumpRules,
		id:              "dependency_php",
	}
}

func (s *PHPDependencySet) ID() string {
	return s.id
}

func (s *PHPDependencySet) SetID(name string) {
	s.id = name
}

func (s *PHPDependencySet) Init() []error {
	composerParser := parser.NewParser(s.pathProvider.PathToCaller())
	composerJson, composerLock, err := composerParser.Parse()
	if err != nil {
		return []error{err}
	}

	cfg := composerConfig.NewComposerConfig(composerJson, composerLock)
	s.bumpDependencies = phpDependenciesToBumpDependencies(cfg.Dependencies)
	s.substituteDependencies = phpDependenciesToSubstituteDependencies(cfg.Dependencies)

	return nil
}

func (s *PHPDependencySet) Run() *rules.RuleTree {
	return rules.NewRuleTree(
		leaf(
			substitute.NewSubstitutePHPLibraryRule(
				s.ID(),
				s.substituteRules,
				s.substituteDependencies,
			),
		),
		leaf(
			bump.NewBumpPHPLibraryRule(
				s.ID(),
				s.bumpRules,
				s.bumpDependencies,
			),
		),
	)
}

func phpDependenciesToSubstituteDependencies(dependencies *composer_json.ComposerDependencies) []substitute.Dependency {
	result := make([]substitute.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}

func phpDependenciesToBumpDependencies(dependencies *composer_json.ComposerDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}
