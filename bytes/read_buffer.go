package bytes

import "io"

// ReadBuffer is a simple read buffer for bytes.
type ReadBuffer struct {
	buf  []byte
	pos  int
	auto bool
}

// MakeReadBuffer creates a new read buffer with the given byte slice.
func MakeReadBuffer(buf []byte, pos int, auto bool) ReadBuffer {
	return ReadBuffer{buf: buf, pos: pos, auto: auto}
}

// NewReadBuffer creates a new read buffer with the given byte slice.
func NewReadBuffer(buf []byte, pos int, auto bool) *ReadBuffer {
	return &ReadBuffer{buf: buf, pos: pos, auto: auto}
}

// Bytes returns the byte slice of buffer.
func (self *ReadBuffer) Bytes() []byte {
	return self.buf
}

// Pos returns the current position of buffer (i.e. total bytes read).
func (self *ReadBuffer) Pos() int {
	return self.pos
}

// Read reads up to len(p) bytes from buffer into p.
// If auto is true and the buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read.
// If auto is false and the buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read and error.
// If there is enough bytes to read, it will read len(p) bytes and return the number of bytes read.
func (self *ReadBuffer) Read(p []byte) (int, error) {
	n := len(p)
	l := self.pos + n
	b := len(self.buf)
	if l <= b {
		i := copy(p, self.buf[self.pos:l])
		self.pos += i
		return n, nil
	}

	n = copy(p, self.buf[self.pos:])
	self.pos += n
	if self.auto {
		return n, nil
	}
	return n, io.ErrShortBuffer
}
