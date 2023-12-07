package sortedmap

// Has checks if the key exists in the collection.
func (sm *SortedMap[K, V]) Has(key K) bool {
	_, ok := sm.Idx[key]
	return ok
}

// BatchHas checks if the keys exist in the collection and returns a slice containing the results.
func (sm *SortedMap[K, V]) BatchHas(keys []K) []bool {
	results := make([]bool, len(keys))
	for i, key := range keys {
		_, results[i] = sm.Idx[key]
	}
	return results
}
