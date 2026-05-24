package main

import (
	"fmt"
	"net/http"
)

var (
	prev int = 0
	next int = 1
)

func HandFib(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, prev)
	prev, next = next, prev+next
}
func main() {
	http.HandleFunc("/", HandFib)
	http.ListenAndServe(":8080", nil)
}
