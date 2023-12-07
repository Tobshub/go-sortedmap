package sortedmap

import "sort"

func (sm *SortedMap[K, V]) setBoundIdx(boundVal V) int {
	return sort.Search(len(sm.sorted), func(i int) bool {
		return sm.lessFn(boundVal, sm.idx[sm.sorted[i]])
	})
}

func (sm *SortedMap[K, V]) boundsIdxSearch(lowerBound, upperBound V) []int {
	smLen := len(sm.sorted)
	if smLen == 0 {
		return nil
	}

	NilV := *new(V)

	if lowerBound != NilV && upperBound != NilV {
		if sm.lessFn(upperBound, lowerBound) {
			return nil
		}
	}

	lowerBoundIdx := 0
	if lowerBound != NilV {
		lowerBoundIdx = sm.setBoundIdx(lowerBound)

		if lowerBoundIdx == smLen {
			lowerBoundIdx--
		}
		if lowerBoundIdx >= 0 && sm.lessFn(sm.idx[sm.sorted[lowerBoundIdx]], lowerBound) {
			lowerBoundIdx++
		}
	}

	upperBoundIdx := smLen - 1
	if upperBound != NilV {
		upperBoundIdx = sm.setBoundIdx(upperBound)
		if upperBoundIdx == smLen {
			upperBoundIdx--
		}
		if upperBoundIdx >= 0 && sm.lessFn(upperBound, sm.idx[sm.sorted[upperBoundIdx]]) {
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
