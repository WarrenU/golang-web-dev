package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	data := []hotel{
		hotel{"Holiday Inn", "1234 Fake Street", "Fresno", "69420", "Southern"},
		hotel{"Good Deal Inn", "1234 Fake Street", "San Jose", "69421", "Central"},
		hotel{"Motel Six", "1234 Fake Street", "Sacramento", "69422", "Northern"},
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
