package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
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
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}
