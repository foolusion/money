package main

import (
	"encoding/gob"
	"html/template"
	"net/http"
	"os"
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
	http.HandleFunc("/add", addHandler)
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
	f, err := os.Open("data.mon")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(&inv)
	if err != nil {
		panic(err)
	}
	page.Invoices = inv
}

var homeTemplate = template.Must(template.ParseFiles("index.html", "head.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := homeTemplate.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("data.mon", os.O_RDWR|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	g := gob.NewEncoder(file)
	g.Encode(inv)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.ListenAndServe(":8080", nil)
}
