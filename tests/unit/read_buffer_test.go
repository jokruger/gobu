package unit

import (
	"testing"

	"github.com/jokruger/gobu"
)

func TestReadBuffer(t *testing.T) {
	bs := []byte("12345")
	buf := gobu.NewReadBufferPtr(bs, 0)
	s, err := fread(buf, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s != "12345" {
		t.Errorf("Unexpected read bytes: %s", s)
	}
	if buf.Pos() != 5 {
		t.Errorf("Unexpected read bytes: %d", buf.Pos())
	}

	bs = []byte("12345")
	buf = gobu.NewReadBufferPtr(bs, 0)
	s, err = fread(buf, 10)
	if err == nil {
		t.Errorf("Expected error")
	}

	bs = []byte("12345")
	buf = gobu.NewReadBufferPtr(bs, 0)
	s, err = fread(buf, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s != "1" {
		t.Errorf("Unexpected read bytes: %s", s)
	}
	s, err = fread(buf, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s != "234" {
		t.Errorf("Unexpected read bytes: %s", s)
	}
	s, err = fread(buf, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s != "5" {
		t.Errorf("Unexpected read bytes: %s", s)
	}
	s, err = fread(buf, 3)
	if err == nil {
		t.Errorf("Expected error")
	}

	bs = []byte("12345")
	buf = gobu.NewReadBufferPtr(bs, 0)
	s, err = fread(buf, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s != "123" {
		t.Errorf("Unexpected read bytes: %s", s)
	}
	s, err = fread(buf, 3)
	if err == nil {
		t.Errorf("Expected error")
	}
}
