package sortedmap

// Get retrieves a value from the collection, using the given key.
func (sm *SortedMap[K, V]) Get(key K) (V, bool) {
	val, ok := sm.idx[key]
	return val, ok
}

// BatchGet retrieves values with their read statuses from the collection, using the given keys.
func (sm *SortedMap[K, V]) BatchGet(keys []K) ([]V, []bool) {
	vals := make([]V, len(keys))
	results := make([]bool, len(keys))

	for i, key := range keys {
		vals[i], results[i] = sm.idx[key]
	}

	return vals, results
}
