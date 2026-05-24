package main

import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {

	res := "curl"
	if method != "GET" {
		res += " -X "
		res += method
	}

	parts := strings.Split(headers, "\n")
	for i := 0; i < len(parts)-1; i++ {
		res += fmt.Sprintf(" -H '%s'", parts[i])
	}
	if body != "" {
		res += " --data '" + body + "'"
	}
	if url != "" {
		res += " " + url
	}
	return res
}
