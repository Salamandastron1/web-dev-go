package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

const (
	contentType  = "Content-Type"
	contentHTML  = "text/html; charset=UTF-8"
	contentPlain = "text/plain; charset=UTF-8"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Host string
		Path string
	}{r.Host, r.URL.Path}

	w.Header().Set(contentType, contentHTML)
	err := tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalln("Template execution error:", err)
	}
}
