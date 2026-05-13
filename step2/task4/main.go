package main

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	fin, _ := os.OpenFile(filename, os.O_WRONLY, 0600)
	defer fin.Close()

	fin.Seek(int64(pos), 0)
	fin.WriteString(val)

}
