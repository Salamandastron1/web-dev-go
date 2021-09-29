package main

import (
	"io"
	"net/http"
)

const (
	contentType  = "Content-Type"
	contentHTML  = "text/html; charset=UTF-8"
	contentPlain = "text/plain; charset=UTF-8"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "moo moo")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "meow meow")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/cow/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
