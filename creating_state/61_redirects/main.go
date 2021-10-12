package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at foo: %s\n\n", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar: %s\n\n", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at barred: %s\n\n", r.Method)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
