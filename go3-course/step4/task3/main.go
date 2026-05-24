package main

import (
	"fmt"
	"net/http"
)

var (
	prev  int = 0
	next  int = 1
	count int = 0
)

func HandFib(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, prev)
	prev, next = next, prev+next
	count++
}
func HandFibReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, count)
}
func main() {
	http.HandleFunc("/", HandFib)
	http.HandleFunc("/metrics", HandFibReq)
	http.ListenAndServe(":8080", nil)
}
