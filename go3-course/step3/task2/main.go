package main

import (
	"strings"
)

func BuildHTTPResponse(statusLine, headers, body string) string {
	var builder strings.Builder
	builder.WriteString(statusLine + "\r\n")
	if headers != "" {
		builder.WriteString(headers)
	}
	builder.WriteString("\r\n")
	if body != "" {
		builder.WriteString(body)
	}
	return builder.String()
}
