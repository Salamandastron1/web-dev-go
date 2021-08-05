package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type restaurant struct {
	Name      string
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	d := restaurant{
		Name:      "Diner",
		Breakfast: []string{"pancakes", "bacon", "toast"},
		Lunch:     []string{"BLT", "Turkey club", "soup"},
		Dinner:    []string{"Steak", "Lentils", "Baba Ganush"},
	}
	sh := restaurant{
		Name:      "Steak House",
		Breakfast: []string{"pancakes", "bacon", "toast"},
		Lunch:     []string{"BLT", "Turkey club", "soup"},
		Dinner:    []string{"Steak", "Lentils", "Baba Ganush"},
	}
	c := restaurant{
		Name:      "Chinese",
		Breakfast: []string{"pancakes", "bacon", "toast"},
		Lunch:     []string{"BLT", "Turkey club", "soup"},
		Dinner:    []string{"Steak", "Lentils", "Baba Ganush"},
	}
	restaurants := []restaurant{d, sh, c}
	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
