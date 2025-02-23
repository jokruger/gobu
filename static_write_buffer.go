package gobu

import "io"

// StaticWriteBuffer is a static write buffer.
type StaticWriteBuffer struct {
	buf []byte
	pos int
}

// NewStaticWriteBuffer creates a new static write buffer with the given byte slice.
func NewStaticWriteBuffer(buf []byte, pos int) StaticWriteBuffer {
	return StaticWriteBuffer{buf: buf, pos: pos}
}

// NewStaticWriteBufferPtr creates a new static write buffer with the given byte slice.
func NewStaticWriteBufferPtr(buf []byte, pos int) *StaticWriteBuffer {
	return &StaticWriteBuffer{buf: buf, pos: pos}
}

// Bytes returns the byte slice of buffer.
func (wb *StaticWriteBuffer) Bytes() []byte {
	return wb.buf
}

// Pos returns the current position of buffer.
func (wb *StaticWriteBuffer) Pos() int {
	return wb.pos
}

// Write writes up to len(p) bytes from p to buffer.
// If buffer does not have enough space to write, it will write available space and return the number of bytes written and error.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (wb *StaticWriteBuffer) Write(p []byte) (int, error) {
	n := len(p)
	b := len(wb.buf)
	if wb.pos+n <= b {
		copy(wb.buf[wb.pos:], p)
		wb.pos += n
		return n, nil
	}

	if wb.pos < b {
		i := copy(wb.buf[wb.pos:], p[:b-wb.pos])
		wb.pos += i
		return i, io.ErrShortWrite
	}

	return 0, io.ErrShortWrite
}
