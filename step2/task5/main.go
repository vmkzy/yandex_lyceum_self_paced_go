package main

import (
	"bufio"
	"errors"
	"os"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	fin, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	var res []string

	scanner := bufio.NewScanner(fin)

	for scanner.Scan() {
		line := scanner.Text()
		s := line[:10]
		date, err := time.Parse("02.01.2006", s)
		if err != nil {
			return nil, err
		}
		if (date.After(start) || date.Equal(start)) && (date.Equal(end) || date.Before(end)) {
			res = append(res, line)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("error")
	}
	return res, nil
}
