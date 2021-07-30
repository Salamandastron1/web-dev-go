package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: true,
	}
	jesus := user{
		Name:  "Jesus",
		Motto: "Love All",
		Admin: false,
	}
	meow := user{
		Name:  "",
		Motto: "Meow Mix Meow Mix plz deliver",
		Admin: true,
	}
	users := []user{buddha, jesus, meow}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
