package benchmark

import (
	"testing"

	"github.com/jokruger/gobu"
)

func BenchmarkReadBuffer(b *testing.B) {
	bs := [12000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewBuffer is part of the benchmark because this is how it is used in the tested use case
		buf := gobu.NewReadBuffer(bs[:], 0)
		for range 1000 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}
