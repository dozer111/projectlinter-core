package composer_lock

type ComposerLockAliases struct {
	items map[string]*RawComposerLockAlias
}

func NewComposerLockAliases(rawItems []*RawComposerLockAlias) *ComposerLockAliases {

	items := make(map[string]*RawComposerLockAlias, len(rawItems))

	for _, alias := range rawItems {
		items[alias.Package] = alias
	}

	return &ComposerLockAliases{
		items: items,
	}
}

func (a *ComposerLockAliases) Has(packageName string) bool {
	_, ok := a.items[packageName]

	return ok
}

func (a *ComposerLockAliases) Get(packageName string) *RawComposerLockAlias {
	return a.items[packageName]
}
