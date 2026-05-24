package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func handHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello stranger")
		return
	}
	if regexp.MustCompile(`[a-zA-Z] + $`).MatchString(name) == true {
		fmt.Fprint(w, "hello dirty hacker")
		return
	}
	fmt.Fprintf(w, "hellp %s", name)

}

func main() {
	http.HandleFunc("/", handHello)
	http.ListenAndServe(":8080", nil)
}
