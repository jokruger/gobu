package benchmark

import (
	"testing"

	"github.com/jokruger/gobu"
)

func BenchmarkReadBuffer(b *testing.B) {
	bs := [1000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := gobu.NewReadBuffer(bs[:], 0)
		for range 84 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}
