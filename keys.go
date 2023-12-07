package sortedmap

import "errors"

func (sm *SortedMap[K, V]) keys(lowerBound, upperBound V) ([]K, error) {
	idxBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if idxBounds == nil {
		return nil, errors.New(noValuesErr)
	}
	return sm.Sorted[idxBounds[0] : idxBounds[1]+1], nil
}

// Keys returns a slice containing sorted keys.
// The returned slice is valid until the next modification to the SortedMap structure.
func (sm *SortedMap[K, V]) Keys() []K {
	keys, _ := sm.keys(*new(V), *new(V))
	return keys
}

// BoundedKeys returns a slice containing sorted keys equal to or between the given bounds.
// The returned slice is valid until the next modification to the SortedMap structure.
func (sm *SortedMap[K, V]) BoundedKeys(lowerBound, upperBound V) ([]K, error) {
	return sm.keys(lowerBound, upperBound)
}
