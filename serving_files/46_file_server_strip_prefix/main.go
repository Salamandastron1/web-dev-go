package main

import (
	"io"
	"log"
	"net/http"
)

const (
	contentType  = "Content-Type"
	contentHTML  = "text/html; charset=UTF-8"
	contentPlain = "text/plain; charset=UTF-8"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentHTML)
	io.WriteString(w, `<!DOCTYPE html><h1>sup</h1><div><image src="/resources/tim.jpeg"></div>`)
}
