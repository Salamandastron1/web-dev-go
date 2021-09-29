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

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(contentType, contentHTML)

	switch req.URL.Path {
	case "/cow":
		io.WriteString(w, "moo moo")
	case "/cat":
		io.WriteString(w, "meow meow")
	}
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}
