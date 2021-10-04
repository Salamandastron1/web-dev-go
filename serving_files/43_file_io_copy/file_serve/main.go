package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	contentType  = "Content-Type"
	contentHTML  = "text/html; charset=UTF-8"
	contentPlain = "text/plain; charset=UTF-8"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/tim.jpeg", timPic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentHTML)
	io.WriteString(w, `
	<img src="/tim.jpeg" />
	`)
}
func timPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("tim.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
