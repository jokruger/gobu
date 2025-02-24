package gobu

import "io"

// ReadBuffer is a read buffer for bytes.
type ReadBuffer struct {
	buf []byte
	off int
}

// NewReadBuffer creates a new read buffer with the given byte slice and offset.
// The caller should not use buf until the buffer is no longer in use.
func NewReadBuffer(buf []byte, off int) *ReadBuffer {
	if buf == nil {
		panic("nil buffer")
	}
	if off < 0 || off > len(buf) {
		panic("invalid offset")
	}
	return &ReadBuffer{buf: buf, off: off}
}

// Buffer returns the underlying byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *ReadBuffer) Buffer() []byte {
	return rb.buf
}

// Processed returns the processed byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *ReadBuffer) Processed() []byte {
	return rb.buf[:rb.off]
}

// Remaining returns the remaining byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *ReadBuffer) Remaining() []byte {
	return rb.buf[rb.off:]
}

// Offset returns the current position in buffer.
func (rb *ReadBuffer) Offset() int {
	return rb.off
}

// Reset resets the buffer offset to zero.
func (rb *ReadBuffer) Reset() {
	rb.off = 0
}

// Read reads up to len(p) bytes from buffer into p.
// If buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read and error.
// If there is enough bytes to read, it will read len(p) bytes and return the number of bytes read.
func (rb *ReadBuffer) Read(p []byte) (int, error) {
	n := len(p)
	l := rb.off + n
	if l <= len(rb.buf) {
		copy(p, rb.buf[rb.off:l])
		rb.off += n
		return n, nil
	}

	if rb.off < len(rb.buf) {
		n = copy(p, rb.buf[rb.off:])
		rb.off += n
		return n, io.EOF
	}

	return 0, io.EOF
}
