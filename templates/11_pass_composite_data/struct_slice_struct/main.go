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
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type items struct {
	Wisdom    []sage
	Transport []car
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
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	c := car{
		Manufacturer: "Toyto",
		Model:        "Corolla",
		Doors:        4,
	}
	sages := []sage{buddha, jesus, meow}
	cars := []car{f, c}
	data := items{
		Wisdom:    sages,
		Transport: cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl_struct_slice_struct.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
