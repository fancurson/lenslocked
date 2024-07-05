package main

import (
	"html/template"
	"os"
)

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets map[string]Widget
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := ViewData{
		Name: "Vlad",
		Widgets: map[string]Widget{
			"a": {Name: "q", Price: 12},
			"b": {Name: "w", Price: 13},
			"c": {Name: "e", Price: 14},
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
