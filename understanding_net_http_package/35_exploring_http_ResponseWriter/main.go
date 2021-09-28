package main

import (
	"fmt"
	"net/http"
)

const (
	contentType = "Content-Type"
	contentHTML = "text/html; charset=UTF-8"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Tim", "sucks.")
	w.Header().Set(contentType, contentHTML)
	fmt.Fprintln(w, "<h1>This is my code, please suck it</h1>")
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}
