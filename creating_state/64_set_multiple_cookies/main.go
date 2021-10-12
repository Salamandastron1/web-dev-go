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
	http.HandleFunc("/abundance", abundance)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie", Value: "meow",
	})
	fmt.Fprintln(w, "Cookie written check your browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	c := r.Cookies()

	fmt.Fprintln(w, "MY COOKIES\n", c)
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "general", Value: "meow",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "specific", Value: "meow",
	})
	fmt.Fprintln(w, "Cookie written check your browser")
}
