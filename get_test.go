package sortedmap

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(3)
	if err != nil {
		t.Fatal(err)
	}

	for i := range records {
		if val, ok := sm.Get(records[i].Key); val == *new(time.Time) || !ok {
			t.Fatalf("TestGet failed: %v", notFoundErr)
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

func TestBatchGet(t *testing.T) {
	sm, _, keys, err := newRandSortedMapWithKeys(1000)
	if err != nil {
		t.Fatal(err)
	}

	values, results := sm.BatchGet(keys)
	for i, ok := range results {
		if values[i] == *new(time.Time) || !ok {
			t.Fatalf("TestBatchGet failed: %v", notFoundErr)
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
