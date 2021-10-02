package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	http.Handle("/dog", http.HandlerFunc(handle))
	http.Handle("/me", http.HandlerFunc(handle))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	var sl []string
	_, err := sb.WriteString("Wow looks like you requested the route: %v\nFrom this host: %v\n")
	if err != nil {
		log.Fatalln(err)
	}
	if r.URL.Path == "/me" {
		sb.WriteString("By the way, my name is Tim")
	}
	sl = strings.Split(fmt.Sprintf(sb.String(), r.URL.Path, r.Host), "\n")
	err = tpl.ExecuteTemplate(w, "index.gohtml", sl)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
}
