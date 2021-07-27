package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl_struct.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
