package package_lock_json

type NPMLockPackages struct {
	items map[string]*RawNPMLockPackage
}

func NewNPMLockPackages(rawItems map[string]*RawNPMLockPackage) *NPMLockPackages {

	items := make(map[string]*RawNPMLockPackage, len(rawItems))

	for packageName, lockPackage := range rawItems {
		items[packageName] = lockPackage
	}

	return &NPMLockPackages{
		items: items,
	}
}

func (p *NPMLockPackages) Has(item string) bool {
	_, ok := p.items[item]

	return ok
}

func (p *NPMLockPackages) Get(item string) *RawNPMLockPackage {
	return p.items[item]
}

func (p *NPMLockPackages) All() map[string]*RawNPMLockPackage {
	return p.items
}

func (p *NPMLockPackages) Count() int {
	return len(p.items)
}
