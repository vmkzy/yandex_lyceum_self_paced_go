package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error) {
	if p == nil {
		return 0, nil
	}
	s := strings.ToUpper(string(p))
	u.UpperString += s
	return len(p), nil
}
