# Overview

[![GoDoc](https://godoc.org/github.com/jokruger/gobu?status.svg)](https://godoc.org/github.com/jokruger/gobu) 
[![Go Report Card](https://goreportcard.com/badge/github.com/jokruger/gobu)](https://goreportcard.com/report/github.com/jokruger/gobu)

Package gobu implements a simple high-performance byte buffers for Go.

## Motivation

The standard library `bytes.Buffer` is a great tool for working with byte slices, but it has some limitations.

**Use Case 1**
Binary serialisation / de-serialisation of data structures where the size of the buffer is known in advance. The `bytes.Buffer` is slower because it does a lot of additional work due to its flexibility and generic implementation.

**Use Case 2**
Custom implementation of the binary encoding/decoding methods where, for compatibility reasons, you need to have a number of popular binary encoding methods like `MarshalBinary`, `GobEncode`, `AppendBinary`, etc., and you want to avoid code duplication.

```go
type MyStruct struct {
    // ...
}

func (s MyStruct) WriteBinary(w io.Writer) error { /* ... */ }

func (s MyStruct) EncodeBinary(buf []byte) (int, error) {
    b := gobu.NewStaticWriteBuffer(buf, 0)
    err := s.WriteBinary(b)
    return b.Offset(), err
}

func (s MyStruct) AppendBinary(buf []byte) ([]byte, error) {
    b := gobu.NewDynamicWriteBuffer(buf, len(buf))
    err := s.WriteBinary(b)
    return b.Buffer(), err
}

func (s MyStruct) MarshalBinary() ([]byte, error) {
    return s.AppendBinary(nil)
}

func (s MyStruct) GobEncode() ([]byte, error) {
    return s.MarshalBinary()
}
```

## Installation

```bash
go get github.com/jokruger/gobu
```

## Benchmarks

```bash
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkReadBytesBuffer-12 (stdlib)            7230 ns/op
BenchmarkReadBytesReader-12 (stdlib)            6630 ns/op
BenchmarkReadBuffer-12 (gobu)                   3511 ns/op

BenchmarkWriteBytesBufferPrealloc-12 (stdlib)   9677 ns/op
BenchmarkStaticWriteBuffer-12 (gobu)            1767 ns/op

BenchmarkWriteBytesBuffer-12 (stdlib)           16308 ns/op
BenchmarkDynamicWriteBuffer-12 (gobu)           12435 ns/op
```
