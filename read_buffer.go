package gobu

import "io"

// ReadBuffer is a read buffer for bytes.
type ReadBuffer struct {
	buf []byte
	pos int
}

// NewReadBuffer creates a new read buffer with the given byte slice.
func NewReadBuffer(buf []byte, pos int) ReadBuffer {
	return ReadBuffer{buf: buf, pos: pos}
}

// NewReadBufferPtr creates a new read buffer with the given byte slice.
func NewReadBufferPtr(buf []byte, pos int) *ReadBuffer {
	return &ReadBuffer{buf: buf, pos: pos}
}

// Bytes returns the byte slice of buffer.
func (rb *ReadBuffer) Bytes() []byte {
	return rb.buf
}

// Pos returns the current position of buffer.
func (rb *ReadBuffer) Pos() int {
	return rb.pos
}

// Read reads up to len(p) bytes from buffer into p.
// If buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read and error.
// If there is enough bytes to read, it will read len(p) bytes and return the number of bytes read.
func (rb *ReadBuffer) Read(p []byte) (int, error) {
	n := len(p)
	l := rb.pos + n
	if l <= len(rb.buf) {
		copy(p, rb.buf[rb.pos:l])
		rb.pos += n
		return n, nil
	}

	if rb.pos < len(rb.buf) {
		n = copy(p, rb.buf[rb.pos:])
		rb.pos += n
		return n, io.ErrShortBuffer
	}

	return 0, io.ErrShortBuffer
}
