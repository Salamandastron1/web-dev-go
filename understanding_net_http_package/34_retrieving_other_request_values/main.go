package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	fmt.Println(data)
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}
