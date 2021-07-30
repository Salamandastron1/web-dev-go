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
	g1 := struct {
		Score1 int
		Score2 int
	}{
		7, 9,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", g1)
	if err != nil {
		log.Fatalln(err)
	}
}
