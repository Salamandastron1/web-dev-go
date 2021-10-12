package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.HandleFunc("/count", count)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	count(w, r)
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie", Value: "meow",
	})

	fmt.Fprintln(w, "Cookie written check your browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	count(w, r)
	c := r.Cookies()

	fmt.Fprintln(w, "MY COOKIES\n", c)
}

func abundance(w http.ResponseWriter, r *http.Request) {
	count(w, r)
	http.SetCookie(w, &http.Cookie{
		Name: "general", Value: "meow",
	})
	http.SetCookie(w, &http.Cookie{
		Name: "specific", Value: "meow",
	})
	fmt.Fprintln(w, "Cookie written check your browser")
}

func count(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("count")
	if err == http.ErrNoCookie {
		log.Println(err, "...Creating")
		c = &http.Cookie{
			Name: "count", Value: "0",
		}
	}
	i, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Invalid numeral. Error: %s", err)
		return
	}
	i++
	c.Value = strconv.Itoa(i)
	http.SetCookie(w, c)
	fmt.Fprintf(w, "This is how many times you have visited: %s\n", c.Value)
}
