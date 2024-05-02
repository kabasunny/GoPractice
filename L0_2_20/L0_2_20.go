package main

import (
	"embed"
	"html/template"
	"log"
	"os"
)

//go:embed template/*
var tmplFS embed.FS

var tmpl = template.Must(template.ParseFS(tmplFS, "template/*.html"))

func main() {
	data := struct {
		Name string
	}{
		Name: "Alice",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "template1.html", data)
	if err != nil {
		log.Fatal(err)
	}
}
