package sortedmap

import (
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/tobshub/go-sortedmap/asc"
)

type (
	TestRecord    = Record[string, time.Time]
	TestSortedMap = SortedMap[string, time.Time]
)

func randStr(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~[]{}|:;<>,./?"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = chars[mrand.Intn(len(chars))]
	}

	return string(result)
}

func randRecord() TestRecord {
	year := mrand.Intn(2129)
	if year < 1 {
		year++
	}
	mth := time.Month(mrand.Intn(12))
	if mth < 1 {
		mth++
	}
	day := mrand.Intn(28)
	if day < 1 {
		day++
	}
	return TestRecord{
		Key: randStr(42),
		Val: time.Date(year, mth, day, 0, 0, 0, 0, time.UTC),
	}
}

func randRecords(n int) []TestRecord {
	records := make([]TestRecord, n)
	for i := range records {
		records[i] = randRecord()
	}
	return records
}

func verifyRecords(ch <-chan TestRecord, reverse bool) error {
	previousRec := TestRecord{}

	if ch != nil {
		for rec := range ch {
			if previousRec.Key != "" {
				switch reverse {
				case false:
					if previousRec.Val.After(rec.Val) {
						return fmt.Errorf("%v %v",
							unsortedErr,
							fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec),
						)
					}
				case true:
					if previousRec.Val.Before(rec.Val) {
						return fmt.Errorf("%v %v",
							unsortedErr,
							fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec),
						)
					}
				}
			}
			previousRec = rec
		}
	} else {
		return fmt.Errorf("Channel was nil.")
	}

	return nil
}

func newSortedMapFromRandRecords(n int) (*TestSortedMap, []TestRecord, error) {
	records := randRecords(n)
	sm := New[string, time.Time](0, asc.Time)
	sm.BatchReplace(records)

	iterCh, err := sm.IterCh()
	if err != nil {
		return sm, records, err
	}
	defer iterCh.Close()

	return sm, records, verifyRecords(iterCh.Records(), false)
}

func newRandSortedMapWithKeys(n int) (*TestSortedMap, []TestRecord, []string, error) {
	sm, records, err := newSortedMapFromRandRecords(n)
	if err != nil {
		return nil, nil, nil, err
	}
	keys := make([]string, n)
	for n, rec := range records {
		keys[n] = rec.Key
	}
	return sm, records, keys, err
}
