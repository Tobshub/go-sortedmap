package sortedmap

import "sort"

func (sm *SortedMap[K, V]) insertSort(key K, val V) []K {
	return insertInterface(sm.sorted, key, sort.Search(len(sm.sorted), func(i int) bool {
		return sm.lessFn(val, sm.idx[sm.sorted[i]])
	}))
}
