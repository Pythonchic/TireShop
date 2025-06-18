package handlers

import (
	"log"
	"net/http"
	"html/template"
)

func New(filename string) (handler http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(filename)
		if err !=  nil {
			log.Printf("Error parse file %s:", filename)
		}
		tmpl.Execute(w, nil)
	}
}
