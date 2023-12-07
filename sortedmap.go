package sortedmap

// SortedMap contains a map, a slice, and references to one or more comparison functions.
// SortedMap is not concurrency-safe, though it can be easily wrapped by a developer-defined type.
type SortedMap[K comparable, V any] struct {
	idx    map[K]V
	sorted []K
	lessFn ComparisonFunc[V]
}

// Record defines a type used in batching and iterations, where keys and values are used together.
type Record[K comparable, V any] struct {
	Key K
	Val V
}

// ComparisonFunc defines the type of the comparison function for the chosen value type.
type ComparisonFunc[V any] func(i, j V) bool

func noOpComparisonFunc[V any](_, _ V) bool {
	return false
}

func setComparisonFunc[V any](cmpFn ComparisonFunc[V]) ComparisonFunc[V] {
	if cmpFn == nil {
		return noOpComparisonFunc[V]
	}
	return cmpFn
}

// New creates and initializes a new SortedMap structure and then returns a reference to it.
// New SortedMaps are created with a backing map/slice of length/capacity n.
func New[K comparable, V any](n int, cmpFn ComparisonFunc[V]) *SortedMap[K, V] {
	return &SortedMap[K, V]{
		idx:    make(map[K]V, n),
		sorted: make([]K, 0, n),
		lessFn: setComparisonFunc(cmpFn),
	}
}

// Len returns the number of items in the collection.
func (sm *SortedMap[K, V]) Len() int {
	return len(sm.sorted)
}
