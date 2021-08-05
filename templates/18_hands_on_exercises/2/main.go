//Create a data structure to pass to a template which
// contains information about California hotels including Name, Address, City, Zip, Region
// region can be: Southern, Central, Northern
// can hold an unlimited number of hotels
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name   string
	Region string
	address
}

type address struct {
	Street string
	City   string
	Zip    int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sb := hotel{
		"Hilton",
		"Southern",
		address{
			"123 Melrose St",
			"Van Nuys",
			81111,
		},
	}
	sf := hotel{
		"Marriot",
		"Northern",
		address{
			"123 Moop St",
			"San Fran",
			000111,
		},
	}
	sd := hotel{
		"Westin",
		"Southern",
		address{
			"123 Poop St",
			"San Diego",
			123445,
		},
	}
	hotels := []hotel{sb, sf, sd}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
