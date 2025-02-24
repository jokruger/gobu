package benchmark

import (
	"testing"

	"github.com/jokruger/gobu"
)

func BenchmarkStaticWriteBuffer(b *testing.B) {
	bs := make([]byte, 0, 12000)
	buf := gobu.NewStaticWriteBuffer(bs, 0)
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewBuffer is not part of the benchmark because in tested use case the buffer can be preallocated
		buf.Reset()
		for range 1000 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}
