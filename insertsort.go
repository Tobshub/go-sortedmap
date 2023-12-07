package sortedmap

import "sort"

func (sm *SortedMap[K, V]) insertSort(key K, val V) []K {
	return insertInterface(sm.Sorted, key, sort.Search(len(sm.Sorted), func(i int) bool {
		return sm.lessFn(val, sm.Idx[sm.Sorted[i]])
	}))
}
