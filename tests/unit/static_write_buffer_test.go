package unit

import (
	"testing"

	"github.com/jokruger/gobu"
)

func TestStaticWriteBuffer(t *testing.T) {
	bs := make([]byte, 10)
	buf := gobu.NewStaticWriteBufferPtr(bs, 0)
	err := fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	written := buf.Pos()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s := string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 5)
	buf = gobu.NewStaticWriteBufferPtr(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	written = buf.Pos()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 4)
	buf = gobu.NewStaticWriteBufferPtr(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err == nil {
		t.Errorf("Expected error")
	}
	written = buf.Pos()
	if written != 4 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "1234" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 3)
	buf = gobu.NewStaticWriteBufferPtr(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err == nil {
		t.Errorf("Expected error")
	}
	written = buf.Pos()
	if written != 3 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "123" {
		t.Errorf("Unexpected written bytes: %s", s)
	}
}
