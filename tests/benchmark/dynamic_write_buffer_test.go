package benchmark

import (
	"testing"

	"github.com/jokruger/gobu"
)

func BenchmarkDynamicWriteBuffer(b *testing.B) {
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewBuffer is part of the benchmark because the goal is to test the allocation strategy
		buf := gobu.NewDynamicWriteBuffer(nil, 0)
		for range 1000 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}
