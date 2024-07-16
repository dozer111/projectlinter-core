package composer_json

type ComposerDependencies struct {
	dependencies map[string]*ComposerDependency
}

func NewComposerDependencies() *ComposerDependencies {
	return &ComposerDependencies{
		dependencies: make(map[string]*ComposerDependency, 30),
	}
}

func (d *ComposerDependencies) Add(dependencies ...*ComposerDependency) {
	for _, dependency := range dependencies {
		d.dependencies[dependency.Name()] = dependency
	}
}

func (d *ComposerDependencies) Has(dependency string) bool {
	_, ok := d.dependencies[dependency]

	return ok
}

func (d *ComposerDependencies) Get(dependency string) *ComposerDependency {
	return d.dependencies[dependency]
}

func (d *ComposerDependencies) All() map[string]*ComposerDependency {
	return d.dependencies
}

func (d *ComposerDependencies) Count() int {
	return len(d.dependencies)
}

func (d *ComposerDependencies) Merge(d2 *ComposerDependencies) *ComposerDependencies {
	newDependencies := &ComposerDependencies{
		make(map[string]*ComposerDependency, d.Count()+d2.Count()),
	}

	for _, mergeWith := range d.All() {
		newDependencies.Add(mergeWith)
	}

	for _, mergeWith := range d2.All() {
		newDependencies.Add(mergeWith)
	}

	return newDependencies
}
