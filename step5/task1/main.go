package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {

	buf := make([]byte, 2048)
	res := make([]byte, 0, 4096)
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
		}
		n, err := r.Read(buf)
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		res = append(res, buf[:n]...)
		if bytes.Contains(res, seq) {
			return true, nil
		}
		if len(res) > len(buf) {
			res = res[len(res)-len(buf)+1:]
		}
	}
}
