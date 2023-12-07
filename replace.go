package sortedmap

import "errors"

func (sm *SortedMap[K, V]) replace(key K, val V) {
	sm.delete(key)
	sm.insert(key, val)
}

// Replace uses the provided 'less than' function to insert sort.
// Even if the key already exists, the value will be inserted.
// Use Insert for the alternative functionality.
func (sm *SortedMap[K, V]) Replace(key K, val V) {
	sm.replace(key, val)
}

// BatchReplace adds all given records to the collection.
// Even if a key already exists, the value will be inserted.
// Use BatchInsert for the alternative functionality.
func (sm *SortedMap[K, V]) BatchReplace(recs []Record[K, V]) {
	for _, rec := range recs {
		sm.replace(rec.Key, rec.Val)
	}
}

// BatchReplaceMap adds all map keys and values to the collection.
// Even if a key already exists, the value will be inserted.
// Use BatchInsertMap for the alternative functionality.
func (sm *SortedMap[K, V]) BatchReplaceMap(v map[K]V) error {
	if v == nil {
		return errors.New("Passed nil map")
	}
	for key, val := range v {
		sm.replace(key, val)
	}
	return nil
}
