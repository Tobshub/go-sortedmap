package sortedmap

import (
	"testing"
	"time"

	"github.com/tobshub/go-sortedmap/asc"
)

func TestReplace(t *testing.T) {
	records := randRecords(3)
	sm := New[string, time.Time](0, asc.Time)

	for i := 0; i < 5; i++ {
		for _, rec := range records {
			sm.Replace(rec.Key, rec.Val)
		}
	}

	iterCh, err := sm.IterCh()
	if err != nil {
		t.Fatal(err)
	} else {
		defer iterCh.Close()

		if err := verifyRecords(iterCh.Records(), false); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBatchReplaceMap(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	m := make(map[string]time.Time, len(records))
	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}
	if err := sm.BatchReplaceMap(m); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplaceMapWithNilType(t *testing.T) {
	if err := New[string, time.Time](0, asc.Time).BatchReplaceMap(nil); err == nil {
		t.Fatal("a nil type was allowed where a supported map type is required.")
	}
}
