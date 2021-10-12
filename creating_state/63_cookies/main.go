package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie", Value: "meow",
	})
	fmt.Fprintln(w, "Cookie written check your browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "MY COOKIE", c)
}
