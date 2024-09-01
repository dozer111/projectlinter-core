package package_json

import "github.com/dozer111/projectlinter-core/rules/javascript/npm/config"

type NPMDependencies struct {
	dependencies map[string][]*config.NPMDependency
}

func NewNPMDependencies(len int) *NPMDependencies {
	if len < 1 {
		len = 30
	}

	return &NPMDependencies{
		dependencies: make(map[string][]*config.NPMDependency, len),
	}
}

func (d *NPMDependencies) Add(dependencies ...*config.NPMDependency) {
	for _, dependency := range dependencies {
		if len(d.dependencies[dependency.Name()]) == 0 {
			d.dependencies[dependency.Name()] = make([]*config.NPMDependency, 0, 2)
		}

		d.dependencies[dependency.Name()] = append(d.dependencies[dependency.Name()], dependency)
	}
}

func (d *NPMDependencies) Has(dependency string) bool {
	_, ok := d.dependencies[dependency]

	return ok
}

func (d *NPMDependencies) All() map[string][]*config.NPMDependency {
	return d.dependencies
}

func (d *NPMDependencies) Count() int {
	return len(d.dependencies)
}

func (d *NPMDependencies) Merge(d2 *NPMDependencies) *NPMDependencies {
	newDependencies := &NPMDependencies{
		make(map[string][]*config.NPMDependency, d.Count()+d2.Count()),
	}

	for _, mergeWith := range d.All() {
		for _, concreteDependency := range mergeWith {
			newDependencies.Add(concreteDependency)
		}
	}

	for _, mergeWith := range d2.All() {
		for _, concreteDependency := range mergeWith {
			newDependencies.Add(concreteDependency)
		}
	}

	return newDependencies
}
