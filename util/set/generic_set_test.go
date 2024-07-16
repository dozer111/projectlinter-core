package utilSet_test

import (
	utilSet "github.com/dozer111/projectlinter-core/util/set"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSet(t *testing.T) {
	// Test creating a set with no initial items
	intSet := utilSet.NewSet[int]()
	require.Equal(t, 0, intSet.Len())

	// Test creating a set with initial items
	strSet := utilSet.NewSet("hello", "world")
	require.Equal(t, 2, strSet.Len())
	require.True(t, strSet.Has("hello"))
	require.True(t, strSet.Has("world"))
}

func TestAdd(t *testing.T) {
	set := utilSet.NewSet[int]()
	set.Add(1, 2, 3)

	require.Equal(t, 3, set.Len())
	require.True(t, set.Has(1))
	require.True(t, set.Has(2))
	require.True(t, set.Has(3))
}

func TestHas(t *testing.T) {
	set := utilSet.NewSet("apple", "banana")

	require.True(t, set.Has("apple"))
	require.False(t, set.Has("orange"))
}
