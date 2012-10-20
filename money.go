package main

import (
	"html/template"
	"net/http"
)

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

type Head struct {
	Selected string
	Options  []string
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
}

var homeTemplate = template.Must(template.ParseFiles("index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var p PageContext
	if err := homeTemplate.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
}
