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
)

type Address struct {
	StreetAddress string
	City          string
	State         string
	ZipCode       string
}

type Wrap struct {
	Data map[string]interface{}
}

func main() {
	files, err := filepath.Glob("templates/*.gohtml")
	if err != nil {
		log.Panic(err)
	}

	// Parse and load the templates.
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Panic(err)
	}

	// Wrap the multiple values into a parameter using map.
	data := map[string]interface{}{
		"Title": "Your Order",
		"Address": Address{
			StreetAddress: "1 Main Street",
			City:          "Springfield",
			State:         "CA",
			ZipCode:       "90405",
		},
		"TotalCost": 50.0,
	}

	wrap := Wrap{Data: data}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render the template.
		if err := tmpl.ExecuteTemplate(w, "home", wrap); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
