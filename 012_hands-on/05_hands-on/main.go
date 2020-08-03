package main

import (
	"log"
	"os"
	"text/template"
)

type dish struct {
	Name  string
	Price float64
}

type menu struct {
	MenuType string
	Dishes   []dish
}

type restaurant struct {
	Name  string
	Menus []menu
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	data := []restaurant{
		restaurant{
			Name: "Johnny's Diner",
			Menus: []menu{
				menu{
					MenuType: "Breakfast",
					Dishes: []dish{
						dish{
							Name:  "egg sandwich",
							Price: 4.69,
						},
						dish{
							Name:  "pancake",
							Price: 1.99,
						},
					},
				},
				menu{
					MenuType: "Lunch",
					Dishes: []dish{
						dish{
							Name:  "italian sandwich",
							Price: 4.49,
						},
						dish{
							Name:  "salad",
							Price: 1.39,
						},
					},
				},
				menu{
					MenuType: "Dinner",
					Dishes: []dish{
						dish{
							Name:  "soup",
							Price: 4.31,
						},
						dish{
							Name:  "meat",
							Price: 21.39,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Smokey's Diner",
			Menus: []menu{
				menu{
					MenuType: "Breakfast",
					Dishes: []dish{
						dish{
							Name:  "sausage sandwich",
							Price: 3.31,
						},
						dish{
							Name:  "waffle",
							Price: 2.39,
						},
					},
				},
				menu{
					MenuType: "Lunch",
					Dishes: []dish{
						dish{
							Name:  "meatball sandwich",
							Price: 8.31,
						},
						dish{
							Name:  "salad",
							Price: 5.39,
						},
					},
				},
				menu{
					MenuType: "Dinner",
					Dishes: []dish{
						dish{
							Name:  "burger",
							Price: 13.31,
						},
						dish{
							Name:  "spam musubi",
							Price: 4.39,
						},
					},
				},
			},
		},
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
