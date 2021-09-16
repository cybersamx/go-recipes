package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	port     = 8000
	tmplDir  = "templates"
	tmplName = "home"
)

func main() {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.gohtml", tmplDir))
	if err != nil {
		log.Panic(err)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, tmplName, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
