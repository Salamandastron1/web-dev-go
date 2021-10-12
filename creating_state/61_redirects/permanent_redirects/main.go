package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at foo: %s\n\n", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar: %s\n\n", r.Method)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
