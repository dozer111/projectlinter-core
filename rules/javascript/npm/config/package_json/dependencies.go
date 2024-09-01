package package_json

type NPMDependencies struct {
	dependencies map[string]*NPMDependency
}

func NewNPMDependencies(len int) *NPMDependencies {
	if len < 1 {
		len = 30
	}

	return &NPMDependencies{
		dependencies: make(map[string]*NPMDependency, len),
	}
}

func (d *NPMDependencies) Add(dependencies ...*NPMDependency) {
	for _, dependency := range dependencies {
		d.dependencies[dependency.Name()] = dependency
	}
}

func (d *NPMDependencies) Has(dependency string) bool {
	_, ok := d.dependencies[dependency]

	return ok
}

func (d *NPMDependencies) All() map[string]*NPMDependency {
	return d.dependencies
}

func (d *NPMDependencies) Count() int {
	return len(d.dependencies)
}

func (d *NPMDependencies) Merge(d2 *NPMDependencies) *NPMDependencies {
	newDependencies := &NPMDependencies{
		make(map[string]*NPMDependency, d.Count()+d2.Count()),
	}

	for _, mergeWith := range d.All() {
		newDependencies.Add(mergeWith)
	}

	for _, mergeWith := range d2.All() {
		newDependencies.Add(mergeWith)
	}

	return newDependencies
}
