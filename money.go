package main

import (
	"html/template"
	"net/http"
	"time"
)

var page PageContext
var inv []*Invoice
var acc []*Account

// Holds the name and balance of an account.
type Account struct {
	Name    string
	Balance float64
}

// Represents a transaction on an account.
type Invoice struct {
	Amount  float64
	Subject string
	Details string
	Tags    []string
	Account
	Time time.Time
}

// The selected values for the table display
type Selected struct {
	Money, Add, Recent, Monthly, Balance, Revenue, Search bool
}

// The Head information for displaying a page
type Head struct {
	Selected
	Options []string
}

// A PageContext hold all the information for displaying a page
type PageContext struct {
	Head
	Invoices []*Invoice
	Tags     []string
}

func init() {
	http.HandleFunc("/", homeHandler)
	http.Handle(
		"/stylesheets/",
		http.StripPrefix(
			"/stylesheets/",
			http.FileServer(http.Dir("stylesheets"))))
	page = PageContext{
		Head: Head{
			Selected: Selected{
				Money:   true,
				Recent:  true,
				Balance: true,
			},
		},
	}
	inv = make([]*Invoice, 0, 100)
	acc = make([]*Account, 0, 5)

	// test data
	acc = append(acc, &Account{
		"Savings",
		10000.00,
	})
	acc = append(acc, &Account{
		"Checking",
		100.00,
	})
	inv = append(inv, &Invoice{
		10.0, "test", "testing", []string{"t1", "t2"}, *acc[1], time.Now()})
	acc[1].Balance -= 10.0
	inv = append(inv, &Invoice{
		20.0, "money", "money things", []string{"t1"}, *acc[1], time.Now()})
	page.Invoices = inv
}

var homeTemplate = template.Must(template.ParseFiles("index.html", "head.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := homeTemplate.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
}
