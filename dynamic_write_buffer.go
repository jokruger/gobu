package gobu

// DynamicWriteBuffer is a dynamic write buffer. It can resize the buffer if there is not enough space to write.
type DynamicWriteBuffer struct {
	buf []byte
	off int
}

// NewDynamicWriteBuffer creates a new write buffer with the given byte slice and offset.
// The caller should not use buf until the buffer is no longer in use.
func NewDynamicWriteBuffer(buf []byte, off int) *DynamicWriteBuffer {
	if buf == nil {
		buf = make([]byte, 0, 8192)
	}
	if off < 0 || off > len(buf) {
		panic("invalid offset")
	}
	return &DynamicWriteBuffer{buf: buf, off: off}
}

// Buffer returns the underlying byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *DynamicWriteBuffer) Buffer() []byte {
	return rb.buf
}

// Processed returns the processed byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *DynamicWriteBuffer) Processed() []byte {
	return rb.buf[:rb.off]
}

// Remaining returns the remaining byte slice.
// The caller should not use the returned slice until the buffer is no longer in use.
func (rb *DynamicWriteBuffer) Remaining() []byte {
	return rb.buf[rb.off:]
}

// Offset returns the current position in buffer.
func (rb *DynamicWriteBuffer) Offset() int {
	return rb.off
}

// Reset resets the buffer offset to zero. It retains the underlying storage for use by future writes.
func (wb *DynamicWriteBuffer) Reset() {
	wb.off = 0
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for another n bytes.
func (wb *DynamicWriteBuffer) Grow(n int) {
	if n <= 0 {
		return
	}

	if wb.off+n <= cap(wb.buf) {
		return
	}

	wb.grow(n)
}

func (wb *DynamicWriteBuffer) grow(n int) {
	b := make([]byte, len(wb.buf), (((n+cap(wb.buf)*2)/8192)+1)*8192)
	copy(b, wb.buf)
	wb.buf = b
}

// Write writes up to len(p) bytes from p to buffer.
// If buffer does not have enough space to write, it will resize the buffer and write p to buffer.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (wb *DynamicWriteBuffer) Write(p []byte) (int, error) {
	n := len(p)
	t := wb.off + n

	if t <= len(wb.buf) {
		copy(wb.buf[wb.off:], p)
		wb.off += n
		return n, nil
	}

	if t <= cap(wb.buf) {
		wb.buf = wb.buf[:t]
		copy(wb.buf[wb.off:], p)
		wb.off = t
		return n, nil
	}

	wb.grow(n)
	wb.buf = wb.buf[:t]
	copy(wb.buf[wb.off:], p)
	wb.off = t
	return n, nil
}
