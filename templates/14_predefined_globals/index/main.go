package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
}
