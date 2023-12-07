package sortedmap

import "fmt"

func (sm *SortedMap[K, V]) insert(key K, val V) bool {
	if _, ok := sm.Idx[key]; !ok {
		sm.Idx[key] = val
		sm.Sorted = sm.insertSort(key, val)
		return true
	}
	return false
}

// Insert uses the provided 'less than' function to insert sort and add the value to the collection and returns a value containing the record's insert status.
// If the key already exists, the value will not be inserted. Use Replace for the alternative functionality.
func (sm *SortedMap[K, V]) Insert(key K, val V) bool {
	return sm.insert(key, val)
}

// BatchInsert adds all given records to the collection and returns a slice containing each record's insert status.
// If a key already exists, the value will not be inserted. Use BatchReplace for the alternative functionality.
func (sm *SortedMap[K, V]) BatchInsert(recs []Record[K, V]) []bool {
	results := make([]bool, len(recs))
	for i, rec := range recs {
		results[i] = sm.insert(rec.Key, rec.Val)
	}
	return results
}

// BatchInsertMap adds all map keys and values to the collection.
// If a key already exists, the value will not be inserted and an error will be returned.
// Use BatchReplaceMap for the alternative functionality.
func (sm *SortedMap[K, V]) BatchInsertMap(v map[K]V) error {
	if v == nil {
		return fmt.Errorf("Passed nili map")
	}
	for key, val := range v {
		if !sm.insert(key, val) {
			return fmt.Errorf("Key already exists: %+v", key)
		}
	}
	return nil
}
