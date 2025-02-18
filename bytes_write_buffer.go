package gobu

import "io"

// BytesWriteBuffer is a simple write buffer for bytes.
type BytesWriteBuffer struct {
	buf  []byte
	pos  int
	auto bool
}

// NewBytesWriteBuffer creates a new write buffer with the given byte slice.
func NewBytesWriteBuffer(buf []byte, pos int, auto bool) *BytesWriteBuffer {
	return &BytesWriteBuffer{buf: buf, pos: pos, auto: auto}
}

// Bytes returns the byte slice of buffer.
func (self *BytesWriteBuffer) Bytes() []byte {
	return self.buf
}

// Pos returns the current position of buffer (i.e. total bytes written).
func (self *BytesWriteBuffer) Pos() int {
	return self.pos
}

// Write writes up to len(p) bytes from p to buffer.
// If auto is true and the buffer does not have enough space to write, it will resize the buffer and write p to buffer.
// If auto is false and the buffer does not have enough space to write, it will write available space and return the number of bytes written and error.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (self *BytesWriteBuffer) Write(p []byte) (int, error) {
	n := len(p)
	l := self.pos + n
	b := len(self.buf)
	if l <= b {
		i := copy(self.buf[self.pos:], p)
		self.pos += i
		return n, nil
	}

	if self.pos < b {
		i := copy(self.buf[self.pos:], p[:b-self.pos])
		if self.auto {
			self.buf = append(self.buf, p[i:]...)
			self.pos = len(self.buf)
			return n, nil
		}
		self.pos += i
		return i, io.ErrShortWrite
	}

	if self.auto {
		self.buf = append(self.buf, p...)
		self.pos = len(self.buf)
		return n, nil
	}

	return 0, io.ErrShortWrite
}
