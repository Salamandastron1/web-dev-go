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
	sages := map[string]string{"India": "Gandhi", "America": "MLK", "Meditate": "Buddha", "Love": "Jesus", "Prophet": "Muhammad"}

	fmt.Println("Key of map not included in template")
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
