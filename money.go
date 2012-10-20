package main

import (
	"html/template"
	"net/http"
)

var page PageContext

type Account struct {
	Name    string
	Balance float64
}

type Invoice struct {
	Amount  float64
	Subject string
	Details string
	Tags    []string
	Account
}

type Selected struct {
	Money, Add, Recent, Monthly, Balance, Revenue, Search bool
}

type Head struct {
	Selected
	Options []string
}

type PageContext struct {
	Head
	Invoices []Invoice
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
