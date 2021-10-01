package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/dog", handle)
	http.HandleFunc("/me", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	sb.WriteString("Wow looks like you requested the route: %v\nFrom this host: %v\n")
	if r.URL.Path == "/me" {
		sb.WriteString("By the way, my name is Tim")
	}
	fmt.Fprintf(w, sb.String(), r.URL.Path, r.Host)
}
