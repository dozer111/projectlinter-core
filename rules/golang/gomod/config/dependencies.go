package config

type GomodDependencies struct {
	dependencies map[string][]*GomodDependency
}

func NewGODependencies(len int) *GomodDependencies {
	if len < 1 {
		len = 30
	}

	return &GomodDependencies{
		dependencies: make(map[string][]*GomodDependency, len),
	}
}

func (d *GomodDependencies) Add(dependencies ...*GomodDependency) *GomodDependencies {
	for _, dependency := range dependencies {
		if len(d.dependencies[dependency.Name()]) == 0 {
			d.dependencies[dependency.Name()] = make([]*GomodDependency, 0, 2)
		}

		d.dependencies[dependency.Name()] = append(d.dependencies[dependency.Name()], dependency)
	}

	return d
}

func (d *GomodDependencies) Has(dependency string) bool {
	_, ok := d.dependencies[dependency]

	return ok
}

func (d *GomodDependencies) All() map[string][]*GomodDependency {
	return d.dependencies
}

func (d *GomodDependencies) Count() int {
	return len(d.dependencies)
}

func (d *GomodDependencies) Merge(d2 *GomodDependencies) *GomodDependencies {
	newDependencies := &GomodDependencies{
		make(map[string][]*GomodDependency, d.Count()+d2.Count()),
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
