package benchmark

import (
	"bytes"
	"testing"
)

func BenchmarkReadBytesBuffer(b *testing.B) {
	bs := [12000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewBuffer is part of the benchmark because this is how it is used in the tested use case
		buf := bytes.NewBuffer(bs[:])
		for range 1000 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}

func BenchmarkReadBytesReader(b *testing.B) {
	bs := [12000]byte{}
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewReader is part of the benchmark because this is how it is used in the tested use case
		buf := bytes.NewReader(bs[:])
		for range 1000 {
			buf.Read(p1[:])
			buf.Read(p2[:])
		}
	}
}

func BenchmarkWriteBytesBufferPrealloc(b *testing.B) {
	bs := make([]byte, 0, 12000)
	buf := bytes.NewBuffer(bs)
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

func BenchmarkWriteBytesBuffer(b *testing.B) {
	p1 := [4]byte{}
	p2 := [8]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// NewBuffer is part of the benchmark because the goal is to test the allocation strategy
		var buf bytes.Buffer
		for range 1000 {
			buf.Write(p1[:])
			buf.Write(p2[:])
		}
	}
}
