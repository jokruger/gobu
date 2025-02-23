package benchmark

import (
	"testing"

	"github.com/jokruger/gobu"
)

func BenchmarkStaticWriteBuffer(b *testing.B) {
	bs := make([]byte, 0, 1008)
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := gobu.NewStaticWriteBuffer(bs, 0)
		for range 84 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}
