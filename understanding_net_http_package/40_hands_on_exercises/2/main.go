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
	http.HandleFunc("/", handle)
	http.HandleFunc("/dog", handle)
	http.HandleFunc("/me", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	var sl []string
	sb.WriteString("Wow looks like you requested the route: %v\nFrom this host: %v\n")
	if r.URL.Path == "/me" {
		sb.WriteString("By the way, my name is Tim")
	}
	sl = strings.Split(fmt.Sprintf(sb.String(), r.URL.Path, r.Host), "\n")
	tpl.ExecuteTemplate(w, "index.gohtml", sl)
}
