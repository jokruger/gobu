package unit

import "io"

func fread(r io.Reader, n int) (string, error) {
	bs := make([]byte, n)
	_, err := r.Read(bs)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func fwrite(w io.Writer, s []string) error {
	for _, v := range s {
		_, err := w.Write([]byte(v))
		if err != nil {
			return err
		}
	}
	return nil
}
