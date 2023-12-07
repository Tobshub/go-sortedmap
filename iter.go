package sortedmap

import (
	"errors"
	"time"
)

// IterChCloser allows records to be read through a channel that is returned by the Records method.
// IterChCloser values should be closed after use using the Close method.
type IterChCloser[K comparable, V any] struct {
	ch       chan Record[K, V]
	canceled chan struct{}
}

// Close cancels a channel-based iteration and causes the sending goroutine to exit.
// Close should be used after an IterChCloser is finished being read from.
func (iterCh *IterChCloser[K, V]) Close() error {
	close(iterCh.canceled)
	return nil
}

// Records returns a channel that records can be read from.
func (iterCh *IterChCloser[K, V]) Records() <-chan Record[K, V] {
	return iterCh.ch
}

// IterChParams contains configurable settings for CustomIterCh.
// SendTimeout is disabled by default, though it should be set to allow
// channel send goroutines to time-out.
// BufSize is set to 1 if its field is set to a lower value.
// LowerBound and UpperBound default to regular iteration when left unset.
type IterChParams[V comparable] struct {
	Reversed               bool
	SendTimeout            time.Duration
	BufSize                int
	LowerBound, UpperBound V
}

// IterCallbackFunc defines the type of function that is passed into an IterFunc method.
// The function is passed a record value argument.
type IterCallbackFunc[K comparable, V any] func(rec Record[K, V]) bool

func setBufSize(bufSize int) int {
	// initialBufSize must be >= 1 or a blocked channel send goroutine may not exit.
	// More info: https://github.com/golang/go/wiki/Timeouts
	const initialBufSize = 1

	if bufSize < initialBufSize {
		return initialBufSize
	}
	return bufSize
}

func (sm *SortedMap[K, V]) recordFromIdx(i int) Record[K, V] {
	rec := Record[K, V]{}
	rec.Key = sm.sorted[i]
	rec.Val = sm.idx[rec.Key]

	return rec
}

func (sm *SortedMap[K, V]) sendRecord(iterCh IterChCloser[K, V], sendTimeout time.Duration, i int) bool {
	if sendTimeout <= time.Duration(0) {
		select {
		case <-iterCh.canceled:
			return false

		case iterCh.ch <- sm.recordFromIdx(i):
			return true
		}
	}

	select {
	case <-iterCh.canceled:
		return false

	case iterCh.ch <- sm.recordFromIdx(i):
		return true

	case <-time.After(sendTimeout):
		return false
	}
}

func (sm *SortedMap[K, V]) iterCh(params IterChParams[V]) (IterChCloser[K, V], error) {
	iterBounds := sm.boundsIdxSearch(params.LowerBound, params.UpperBound)
	if iterBounds == nil {
		return IterChCloser[K, V]{}, errors.New(noValuesErr)
	}

	iterCh := IterChCloser[K, V]{
		ch:       make(chan Record[K, V], setBufSize(params.BufSize)),
		canceled: make(chan struct{}),
	}

	go func(params IterChParams[V], iterCh IterChCloser[K, V]) {
		if params.Reversed {
			for i := iterBounds[1]; i >= iterBounds[0]; i-- {
				if !sm.sendRecord(iterCh, params.SendTimeout, i) {
					break
				}
			}
		} else {
			for i := iterBounds[0]; i <= iterBounds[1]; i++ {
				if !sm.sendRecord(iterCh, params.SendTimeout, i) {
					break
				}
			}
		}
		close(iterCh.ch)
	}(params, iterCh)

	return iterCh, nil
}

func (sm *SortedMap[K, V]) iterFunc(reversed bool, lowerBound, upperBound V, f IterCallbackFunc[K, V]) error {
	iterBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if iterBounds == nil {
		return errors.New(noValuesErr)
	}

	if reversed {
		for i := iterBounds[1]; i >= iterBounds[0]; i-- {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	} else {
		for i := iterBounds[0]; i <= iterBounds[1]; i++ {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	}

	return nil
}

// IterCh returns a channel that sorted records can be read from and processed.
// This method defaults to the expected behavior of blocking until a read, with no timeout.
func (sm *SortedMap[K, V]) IterCh() (IterChCloser[K, V], error) {
	return sm.iterCh(IterChParams[V]{})
}

// BoundedIterCh returns a channel that sorted records can be read from and processed.
// BoundedIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap[K, V]) BoundedIterCh(reversed bool, lowerBound, upperBound V) (IterChCloser[K, V], error) {
	return sm.iterCh(IterChParams[V]{
		Reversed:   reversed,
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
}

// CustomIterCh returns a channel that sorted records can be read from and processed.
// CustomIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap[K, V]) CustomIterCh(params IterChParams[V]) (IterChCloser[K, V], error) {
	return sm.iterCh(params)
}

// IterFunc passes each record to the specified callback function.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap[K, V]) IterFunc(reversed bool, f IterCallbackFunc[K, V]) {
	sm.iterFunc(reversed, *new(V), *new(V), f)
}

// BoundedIterFunc starts at the lower bound value and passes all values in the collection to the callback function until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap[K, V]) BoundedIterFunc(reversed bool, lowerBound, upperBound V, f IterCallbackFunc[K, V]) error {
	return sm.iterFunc(reversed, lowerBound, upperBound, f)
}
