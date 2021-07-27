package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	fmt.Println("Simple iteration of data")
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)

	}

	fmt.Println("Select specific Template, variablize data in template")
	err = tpl.ExecuteTemplate(os.Stdout, "tpl_var.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
