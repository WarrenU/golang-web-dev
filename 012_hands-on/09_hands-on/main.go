package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Record is the data recorded from the table.csv file.
type Record struct {
	Date time.Time
	Open float64
}

func handler(res http.ResponseWriter, req *http.Request) {
	records := parse("table.csv")

	template, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = template.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

// parses table.csv and returns a json object of all the data.
func parse(filename string) []Record {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(data))

	for i, row := range data {
		// Skip csv Headers:
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{date, open})
	}

	return records
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
