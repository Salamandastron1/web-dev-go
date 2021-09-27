package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}
