package dependency

import (
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config"
	jsNPMParser "github.com/dozer111/projectlinter-core/rules/javascript/npm/parser"
	"github.com/dozer111/projectlinter-core/util/path_provider"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
)

// JavascriptNPMDependencySet reusable already created RuleSet
//
// # It uses already generated bump configs to check
//
// As for me - the code of dependency module(dependencySet) would always the same
// So no need to write some custom - use this instead
type JavascriptNPMDependencySet struct {
	pathProvider *path_provider.PathProvider

	substituteRules        []substitute.Library
	substituteDependencies []substitute.Dependency

	bumpRules        []bump.Library
	bumpDependencies []bump.Dependency

	id string
}

var _ rules.Set = (*JavascriptNPMDependencySet)(nil)

func NewJavascriptNPMDependencySet(
	pathProvider *path_provider.PathProvider,
	substituteRules []substitute.Library,
	bumpRules []bump.Library,
) *JavascriptNPMDependencySet {
	return &JavascriptNPMDependencySet{
		pathProvider:    pathProvider,
		substituteRules: substituteRules,
		bumpRules:       bumpRules,
		id:              "dependency",
	}
}

func (s *JavascriptNPMDependencySet) ID() string {
	return s.id
}

func (s *JavascriptNPMDependencySet) SetID(name string) {
	s.id = name
}

func (s *JavascriptNPMDependencySet) Init() []error {
	npmParser := jsNPMParser.NewParser(s.pathProvider.PathToCaller())
	packageJSON, packageLock, err := npmParser.Parse()
	if err != nil {
		return []error{err}
	}

	cfg := config.NewNPMConfig(packageJSON, packageLock)
	s.bumpDependencies = javascriptNPMDependenciesToBumpDependencies(cfg.Dependencies)
	s.substituteDependencies = javascriptNPMDependenciesToSubstituteDependencies(cfg.Dependencies)

	return nil
}

func (s *JavascriptNPMDependencySet) Run() *rules.RuleTree {
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

func javascriptNPMDependenciesToSubstituteDependencies(dependencies *config.NPMDependencies) []substitute.Dependency {
	result := make([]substitute.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}

func javascriptNPMDependenciesToBumpDependencies(dependencies *config.NPMDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}
