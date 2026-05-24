package main

import (
	"fmt"
	"strings"
)

func BuildHTTPRequest(method, url, host, headers, body string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s %s HTTP/1.1\r\n", method, url))
	builder.WriteString(fmt.Sprintf("Host: %s\r\n", host))
	if headers != "" {
		builder.WriteString(headers)
	}
	builder.WriteString("\r\n")
	if body != "" {
		builder.WriteString(body)
	}
	return builder.String()
}
