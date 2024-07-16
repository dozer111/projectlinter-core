package composer_lock

type ComposerLockPackages struct {
	items map[string]*RawComposerLockPackage
}

func NewComposerLockPackages(rawItems []*RawComposerLockPackage) *ComposerLockPackages {

	items := make(map[string]*RawComposerLockPackage, len(rawItems))

	for _, lockPackage := range rawItems {
		items[lockPackage.Name] = lockPackage
	}

	return &ComposerLockPackages{
		items: items,
	}
}

func (p *ComposerLockPackages) Has(item string) bool {
	_, ok := p.items[item]

	return ok
}

func (p *ComposerLockPackages) Get(item string) *RawComposerLockPackage {
	return p.items[item]
}

func (p *ComposerLockPackages) All() map[string]*RawComposerLockPackage {
	return p.items
}

func (p *ComposerLockPackages) Count() int {
	return len(p.items)
}
