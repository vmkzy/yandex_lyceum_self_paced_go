package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer f.Close()

	stdScan := bufio.NewScanner(f)
	indx := 0
	for stdScan.Scan() {
		if indx == lineNum {
			return stdScan.Text()
		}
		indx++
	}
	if err := stdScan.Err(); err != nil {
		return ""
	}
	return ""
}
