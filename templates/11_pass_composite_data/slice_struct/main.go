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
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love All",
	}
	meow := sage{
		Name:  "CatFace",
		Motto: "Meow Mix Meow Mix plz deliver",
	}
	sages := []sage{buddha, jesus, meow}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl_slice_struct.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
