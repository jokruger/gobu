package benchmark

import (
	"bytes"
	"testing"
)

func BenchmarkReadBytesBuffer(b *testing.B) {
	bs := [1000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(bs[:])
		for range 84 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}

func BenchmarkReadBytesReader(b *testing.B) {
	bs := [1000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(bs[:])
		for range 84 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}

func BenchmarkWriteBytesBufferPrealloc(b *testing.B) {
	bs := make([]byte, 0, 1008)
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(bs)
		for range 84 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}

func BenchmarkWriteBytesBuffer(b *testing.B) {
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for range 84 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}
