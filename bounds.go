package sortedmap

import (
	"reflect"
	"sort"
)

func (sm *SortedMap[K, V]) setBoundIdx(boundVal V) int {
	return sort.Search(len(sm.Sorted), func(i int) bool {
		return sm.lessFn(boundVal, sm.Idx[sm.Sorted[i]])
	})
}

func (sm *SortedMap[K, V]) boundsIdxSearch(lowerBound, upperBound V) []int {
	smLen := len(sm.Sorted)
	if smLen == 0 {
		return nil
	}

	if !reflect.ValueOf(lowerBound).IsZero() && !reflect.ValueOf(upperBound).IsZero() {
		if sm.lessFn(upperBound, lowerBound) {
			return nil
		}
	}

	lowerBoundIdx := 0
	if !reflect.ValueOf(lowerBound).IsZero() {
		lowerBoundIdx = sm.setBoundIdx(lowerBound)

		if lowerBoundIdx == smLen {
			lowerBoundIdx--
		}
		if lowerBoundIdx >= 0 && sm.lessFn(sm.Idx[sm.Sorted[lowerBoundIdx]], lowerBound) {
			lowerBoundIdx++
		}
	}

	upperBoundIdx := smLen - 1
	if !reflect.ValueOf(upperBound).IsZero() {
		upperBoundIdx = sm.setBoundIdx(upperBound)
		if upperBoundIdx == smLen {
			upperBoundIdx--
		}
		if upperBoundIdx >= 0 && sm.lessFn(upperBound, sm.Idx[sm.Sorted[upperBoundIdx]]) {
			upperBoundIdx--
		}
	}

	if lowerBoundIdx > upperBoundIdx {
		return nil
	}

	return []int{
		lowerBoundIdx,
		upperBoundIdx,
	}
}
