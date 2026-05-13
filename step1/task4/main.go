package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	readData := io.LimitReader(r, int64(n))
	_, err := io.Copy(w, readData)
	if err != nil {
		return err
	}
	return nil
}
