package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wow looks like you requested the default route: %v\nFrom this host: %v", r.URL.Path, r.Host)
	})
	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wow looks like you requested the route: %v\nFrom this host %v", r.URL.Path, r.Host)
	})
	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wow looks like you requested the route: %v\nFrom this host: %v\nBy the way, my name is Tim", r.URL.Path, r.Host)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
