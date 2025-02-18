package gobu

import (
	"io"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	fwrite := func(w io.Writer, s []string) error {
		for _, v := range s {
			_, err := w.Write([]byte(v))
			if err != nil {
				return err
			}
		}
		return nil
	}

	t.Run("Fixed writer", func(t *testing.T) {
		bs := make([]byte, 10)
		buf := NewBytesWriteBuffer(bs, 0, false)
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
		buf = NewBytesWriteBuffer(bs, 0, false)
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
		buf = NewBytesWriteBuffer(bs, 0, false)
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
		buf = NewBytesWriteBuffer(bs, 0, false)
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
	})

	t.Run("Auto writer", func(t *testing.T) {
		var bs []byte
		buf := NewBytesWriteBuffer(bs, 0, true)
		err := fwrite(buf, []string{"1", "234", "5"})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		bs = buf.Bytes()
		written := buf.Pos()
		if written != 5 {
			t.Errorf("Unexpected written bytes: %d", written)
		}
		s := string(bs[:written])
		if s != "12345" {
			t.Errorf("Unexpected written bytes: %s", s)
		}

		bs = make([]byte, 10)
		buf = NewBytesWriteBuffer(bs, 0, true)
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

		bs = make([]byte, 5)
		buf = NewBytesWriteBuffer(bs, 0, true)
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
		buf = NewBytesWriteBuffer(bs, 0, true)
		err = fwrite(buf, []string{"1", "234", "5"})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		bs = buf.Bytes()
		written = buf.Pos()
		if written != 5 {
			t.Errorf("Unexpected written bytes: %d", written)
		}
		s = string(bs[:written])
		if s != "12345" {
			t.Errorf("Unexpected written bytes: %s", s)
		}

		bs = make([]byte, 3)
		buf = NewBytesWriteBuffer(bs, 0, true)
		err = fwrite(buf, []string{"1", "234", "5"})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		bs = buf.Bytes()
		written = buf.Pos()
		if written != 5 {
			t.Errorf("Unexpected written bytes: %d", written)
		}
		s = string(bs[:written])
		if s != "12345" {
			t.Errorf("Unexpected written bytes: %s", s)
		}
	})

	fread := func(r io.Reader, n int) (string, error) {
		bs := make([]byte, n)
		_, err := r.Read(bs)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}

	t.Run("Fixed reader", func(t *testing.T) {
		bs := []byte("12345")
		buf := NewBytesReadBuffer(bs, 0, false)
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
		buf = NewBytesReadBuffer(bs, 0, false)
		s, err = fread(buf, 10)
		if err == nil {
			t.Errorf("Expected error")
		}

		bs = []byte("12345")
		buf = NewBytesReadBuffer(bs, 0, false)
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
		buf = NewBytesReadBuffer(bs, 0, false)
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
	})
}
