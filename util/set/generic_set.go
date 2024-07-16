package utilSet

// Set represents a generic set
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new set
func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}
	s.Add(items...)

	return s
}

// Add adds an item to the set
func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

// Has checks if the item exists in the set
func (s *Set[T]) Has(item T) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
