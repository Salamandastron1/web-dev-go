package main

import (
	"encoding/csv"
	"log"
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

		date, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
}
