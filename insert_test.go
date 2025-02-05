package sortedmap

import (
	"testing"
	"time"

	"github.com/tobshub/go-sortedmap/asc"
)

func TestInsert(t *testing.T) {
	const n = 3
	records := randRecords(n)
	sm := New[string, time.Time](n, asc.Time)

	for i := range records {
		if !sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", keyExistsErr)
		}
	}

	func() {
		iterCh, err := sm.IterCh()
		if err != nil {
			t.Fatal(err)
		} else {
			defer iterCh.Close()

			if err := verifyRecords(iterCh.Records(), false); err != nil {
				t.Fatal(err)
			}
		}
	}()

	for i := range records {
		if sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", notFoundErr)
		}
	}

	func() {
		iterCh, err := sm.IterCh()
		if err != nil {
			t.Fatal(err)
		} else {
			defer iterCh.Close()

			if err := verifyRecords(iterCh.Records(), false); err != nil {
				t.Fatal(err)
			}
		}
	}()
}

func TestBatchInsert(t *testing.T) {
	const n = 1000
	records := randRecords(n)
	sm := New[string, time.Time](n, asc.Time)

	for _, ok := range sm.BatchInsert(records) {
		if !ok {
			t.Fatalf("BatchInsert failed: %v", keyExistsErr)
		}
	}

	func() {
		iterCh, err := sm.IterCh()
		if err != nil {
			t.Fatal(err)
		} else {
			defer iterCh.Close()

			if err := verifyRecords(iterCh.Records(), false); err != nil {
				t.Fatal(err)
			}
		}
	}()
}

func TestBatchInsertMap(t *testing.T) {
	const n = 1000
	records := randRecords(n)
	sm := New[string, time.Time](n, asc.Time)

	i := 0
	m := make(map[string]time.Time, n)

	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}

	if err := sm.BatchInsertMap(m); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsertMapWithExistingInterfaceKeys(t *testing.T) {
	const n = 1000
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	i := 0
	m := make(map[string]time.Time, n)

	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}

	if err := sm.BatchInsertMap(m); err == nil {
		t.Fatal("Inserting existing keys should have caused an error.")
	}
}

func TestBatchInsertMapWithExistingStringKeys(t *testing.T) {
	const n = 1000
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	i := 0
	m := make(map[string]time.Time, n)

	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}

	if err := sm.BatchInsertMap(m); err == nil {
		t.Fatal("Inserting existing keys should have caused an error.")
	}
}

func TestBatchInsertMapWithNilType(t *testing.T) {
	if err := New[string, time.Time](0, asc.Time).BatchInsertMap(nil); err == nil {
		t.Fatal("a nil type was allowed where a supported map type is required.")
	}
}
