package main

import (
	"bytes"
	"context"
	"io"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)

	if ctx.Err() != nil {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	var data bytes.Buffer
	buffer := make([]byte, 32*1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		count, err := file.Read(buffer)
		if count > 0 {
			if _, writeErr := data.Write(buffer[:count]); writeErr != nil {
				return
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
	}

	select {
	case <-ctx.Done():
		return
	case result <- data.Bytes():
	}
}
