package main

import (
	"log"
	"os"
	"text/template"
)

var tplte *template.Template

func init() { // load templates in memory, run once
	tplte = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// petNames := []string{"Ellie", "Novee"}

	petNames := map[string]string{
		"Ellie": "Dog",
		"Novee": "Doggie",
	}

	err := tplte.Execute(os.Stdout, petNames)

	if err != nil {
		log.Fatalln(err)
	}
}


