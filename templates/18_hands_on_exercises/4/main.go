package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

type Record struct {
	Date time.Time
	Open float64
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", generateTemplate)
	http.ListenAndServe(":8080", nil)
}

func generateTemplate(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	records := make([]Record, 0, len(lines))
	for i, row := range lines {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}
