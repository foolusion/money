package main

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
}

var homeTemplate = template.Must(template.ParseFiles("index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := homeTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
}
