package main

import (
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {

	buf := make([]byte, 0, 4096+len(seq))
	tmp := make([]byte, 4096)

	for {
		n, err := r.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
			if idx := bytes.Index(buf, seq); idx != -1 {
				return true, nil
			}
			if len(buf) > len(seq)-1 {
				start := len(buf) - (len(seq) - 1)
				if start < 0 {
					start = 0
				}
				tail := make([]byte, len(buf)-start)
				copy(tail, buf[start:])
				buf = tail
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}
	}

	return false, nil
}
