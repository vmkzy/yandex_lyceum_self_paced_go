package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	fin, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer fin.Close()

	fout, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = fin.Seek(int64(startpos), 0)

	if err != nil {
		return err
	}
	_, err = io.Copy(fout, fin)
	if err != nil {
		return err
	}
	return nil
}
