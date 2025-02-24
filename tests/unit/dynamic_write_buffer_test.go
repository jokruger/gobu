package unit

import (
	"testing"

	"github.com/jokruger/gobu"
)

func TestDynamicWriteBuffer(t *testing.T) {
	var bs []byte
	buf := gobu.NewDynamicWriteBuffer(bs, 0)
	err := fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	bs = buf.Buffer()
	written := buf.Offset()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s := string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 10)
	buf = gobu.NewDynamicWriteBuffer(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	written = buf.Offset()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 5)
	buf = gobu.NewDynamicWriteBuffer(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	written = buf.Offset()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 4)
	buf = gobu.NewDynamicWriteBuffer(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	bs = buf.Buffer()
	written = buf.Offset()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}

	bs = make([]byte, 3)
	buf = gobu.NewDynamicWriteBuffer(bs, 0)
	err = fwrite(buf, []string{"1", "234", "5"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	bs = buf.Buffer()
	written = buf.Offset()
	if written != 5 {
		t.Errorf("Unexpected written bytes: %d", written)
	}
	s = string(bs[:written])
	if s != "12345" {
		t.Errorf("Unexpected written bytes: %s", s)
	}
}
