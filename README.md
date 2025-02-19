# Overview

[![GoDoc](https://godoc.org/github.com/jokruger/gobu?status.svg)](https://godoc.org/github.com/jokruger/gobu) 
[![Go Report Card](https://goreportcard.com/badge/github.com/jokruger/gobu)](https://goreportcard.com/report/github.com/jokruger/gobu)

Package gobu is a collection of interfaces and helper functions to simplify the implementation of the `encoding.BinaryMarshaler`, `encoding.BinaryUnmarshaler`, `encoding.BinaryAppender` and `encoding.GobEncoder` interfaces.

## Usage Example

```go
package mypackage

import (
	"io"
	"github.com/jokruger/gobu/bytes"
)

type MyStruct struct {
    // your fields here
}

func (m MyStruct) WriteBinary(w io.Writer) error { /* your implementation here */ }

func (m MyStruct) AppendBinary(buf []byte) ([]byte, error) {
	b := bytes.NewWriteBuffer(buf, len(buf), true)
	err := m.WriteBinary(b)
	return b.Bytes(), err
}

func (m MyStruct) MarshalBinary() ([]byte, error) {
	return m.AppendBinary(nil)
}

func (m MyStruct) GobEncode() ([]byte, error) {
	return m.MarshalBinary()
}
```

In this example the `MyStruct` type implements the `WriteBinary` method which writes the struct to a `io.Writer`. Additionally it implements the `encoding.BinaryAppender`, `encoding.BinaryMarshaler` and `encoding.GobEncoder` interfaces which are using the `WriteBinary` method to write the struct to a buffer.

## Installation

```bash
go get github.com/jokruger/gobu
```
