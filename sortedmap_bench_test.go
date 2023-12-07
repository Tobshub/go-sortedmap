package sortedmap

import (
	"testing"
	"time"

	"github.com/tobshub/go-sortedmap/asc"
)

func BenchmarkNew(b *testing.B) {
	var sm *TestSortedMap

	for i := 0; i < b.N; i++ {
		sm = New[string, time.Time](0, asc.Time)
	}

	if sm == nil {
	}
}
