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
	// if http.ListenAndServe exits it captures the error
	// then prints it to the log for later debugging
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentHTML)
	io.WriteString(w, `<img src="/tim.jpeg" />`)
}
func timPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("tim.jpeg")
	if err != nil {
		//write to the network connection with error
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
