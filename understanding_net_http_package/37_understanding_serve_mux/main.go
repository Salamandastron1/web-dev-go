package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "moo moo")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "meow meow")
}

func main() {
	http.HandleFunc("/cow", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
