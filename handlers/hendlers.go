package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func LoginTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/authorisation.page.tmpl")
	if err != nil {
		log.Println("Error to parse file authorization.page.tmpl", err)
	}
	t.Execute(w, nil)
}
