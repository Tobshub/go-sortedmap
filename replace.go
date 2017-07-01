package sortedmap

func (sm *SortedMap) replace(key, val interface{}) {
	sm.delete(key)
	sm.insert(key, val)
}

// Replace uses the provided 'less than' function to insert sort.
// Even if the key already exists, the value will be inserted. Use Insert for the alternative functionality.
func (sm *SortedMap) Replace(key, val interface{}) {
	sm.replace(key, val)
}

// BatchReplace adds all given records to the collection.
// Even if a key already exists, the value will be inserted. Use BatchInsert for the alternative functionality.	
func (sm *SortedMap) BatchReplace(recs []*Record) {
	for _, rec := range recs {
		sm.replace(rec.Key, rec.Val)
	}
}

// BatchReplaceMap adds all map keys and values to the collection.
// Even if a key already exists, the value will be inserted. Use BatchInsertMap for the alternative functionality.	
func (sm *SortedMap) BatchReplaceMap(v interface{}) bool {
	switch m := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range m {
			sm.replace(key, val)
		}
		return true

	case map[string]interface{}:
		for key, val := range m {
			sm.replace(key, val)
		}
		return true
	}
	return false
}