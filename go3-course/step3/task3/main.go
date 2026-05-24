package main

import (
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string) {
	parts := strings.SplitN(statusLine, " ", 3)
	code, _ = strconv.Atoi(parts[1])
	reason = parts[2]
	return code, reason
}
