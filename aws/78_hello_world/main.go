package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":80", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}
