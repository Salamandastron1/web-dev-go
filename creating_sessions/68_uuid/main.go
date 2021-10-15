package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}
	fmt.Println(c)
}
