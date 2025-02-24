package gobu

import "io"

// StaticWriteBuffer is a static write buffer.
type StaticWriteBuffer struct {
	buf []byte
	off int
}

// NewStaticWriteBuffer creates a new static write buffer with the given byte slice and offset.
// The caller should not use buf until the buffer is no longer in use.
func NewStaticWriteBuffer(buf []byte, off int) *StaticWriteBuffer {
	if buf == nil {
		panic("nil buffer")
	}
	if off < 0 || off > len(buf) {
		panic("invalid offset")
	}
	return &StaticWriteBuffer{buf: buf, off: off}
}

// Buffer returns the underlying byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *StaticWriteBuffer) Buffer() []byte {
	return rb.buf
}

// Processed returns the processed byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *StaticWriteBuffer) Processed() []byte {
	return rb.buf[:rb.off]
}

// Remaining returns the remaining byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *StaticWriteBuffer) Remaining() []byte {
	return rb.buf[rb.off:]
}

// Offset returns the current position in buffer.
func (rb *StaticWriteBuffer) Offset() int {
	return rb.off
}

// Reset resets the buffer offset to zero.
func (wb *StaticWriteBuffer) Reset() {
	wb.off = 0
}

// Write writes up to len(p) bytes from p to buffer.
// If buffer does not have enough space to write, it will write available space and return the number of bytes written and error.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (wb *StaticWriteBuffer) Write(p []byte) (int, error) {
	n := len(p)
	b := len(wb.buf)
	if wb.off+n <= b {
		copy(wb.buf[wb.off:], p)
		wb.off += n
		return n, nil
	}

	if wb.off < b {
		i := copy(wb.buf[wb.off:], p[:b-wb.off])
		wb.off += i
		return i, io.ErrShortWrite
	}

	return 0, io.ErrShortWrite
}
