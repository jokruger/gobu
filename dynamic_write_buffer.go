package gobu

// DynamicWriteBuffer is a dynamic write buffer. It can resize the buffer if there is not enough space to write.
type DynamicWriteBuffer struct {
	buf []byte
	pos int
}

// NewDynamicWriteBuffer creates a new write buffer with the given byte slice.
func NewDynamicWriteBuffer(buf []byte, pos int) DynamicWriteBuffer {
	return DynamicWriteBuffer{buf: buf, pos: pos}
}

// NewDynamicWriteBufferPtr creates a new write buffer with the given byte slice.
func NewDynamicWriteBufferPtr(buf []byte, pos int) *DynamicWriteBuffer {
	return &DynamicWriteBuffer{buf: buf, pos: pos}
}

// Bytes returns the byte slice of buffer.
func (wb *DynamicWriteBuffer) Bytes() []byte {
	return wb.buf
}

// Pos returns the current position of buffer.
func (wb *DynamicWriteBuffer) Pos() int {
	return wb.pos
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for another n bytes.
func (wb *DynamicWriteBuffer) Grow(n int) {
	if n <= 0 {
		return
	}

	if wb.pos+n <= cap(wb.buf) {
		return
	}

	wb.grow(n)
}

func (wb *DynamicWriteBuffer) grow(n int) {
	b := make([]byte, len(wb.buf), max(wb.pos+n, cap(wb.buf)*2, 2048))
	copy(b, wb.buf)
	wb.buf = b
}

// Write writes up to len(p) bytes from p to buffer.
// If buffer does not have enough space to write, it will resize the buffer and write p to buffer.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (wb *DynamicWriteBuffer) Write(p []byte) (int, error) {
	n := len(p)
	t := wb.pos + n

	if t <= len(wb.buf) {
		copy(wb.buf[wb.pos:], p)
		wb.pos += n
		return n, nil
	}

	if t <= cap(wb.buf) {
		wb.buf = wb.buf[:t]
		copy(wb.buf[wb.pos:], p)
		wb.pos = t
		return n, nil
	}

	wb.grow(n)
	wb.buf = wb.buf[:t]
	copy(wb.buf[wb.pos:], p)
	wb.pos = t
	return n, nil
}
