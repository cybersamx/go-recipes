package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	port         = 8000
	templateFile = "templates/home.html"
)

type Address struct {
	StreetAddress string
	City          string
	State         string
	ZipCode       string
}

type Customer struct {
	FirstName string
	LastName  string
	Email     string
}

type Item struct {
	Product string
	Qty     int
	Cost    float64
}

type Order struct {
	Shipper    string
	Customer   *Customer
	ShipmentID string
	Address    *Address
	Items      []*Item
}

func (c *Customer) FullName(lastNameFirst bool) string {
	if lastNameFirst {
		return fmt.Sprintf("%s, %s", c.LastName, c.FirstName)
	}

	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (o *Order) TotalCost() float64 {
	var totCost float64
	for _, item := range o.Items {
		totCost += float64(item.Qty) * item.Cost
	}
	return totCost
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(templateFile))
		htmlContent := &Order{
			Shipper: "Acme",
			Customer: &Customer{
				FirstName: "Kirby",
				LastName:  "Bryan",
				Email:     "mvp@example.com",
			},
			ShipmentID: "ABC-XY-123",
			Address: &Address{
				StreetAddress: "1111 S Figueroa St",
				City:          "Los Angeles",
				State:         "CA",
				ZipCode:       "90015",
			},
			Items: []*Item{
				{Product: "Basket Ball", Qty: 1, Cost: 30.0},
				{Product: "Sneakers", Qty: 1, Cost: 90.5},
				{Product: "Sock", Qty: 2, Cost: 15.0},
			},
		}
		if err := tmpl.Execute(w, htmlContent); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func main() {
	http.Handle("/", homeHandler())
	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
