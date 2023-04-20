package handlers

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"html/template"
	"log"
	"net/http"
	"time"
)

func LoginTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/authorization.page.tmpl")
	if err != nil {
		log.Println("Error to parse file authorization.page.tmpl", err)
	}
	t.Execute(w, nil)
}

func ChatGPT(tokens map[string]authorization.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenValue := r.Header.Get("Authorization")
		token, ok := tokens[tokenValue]
		if !ok || token.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		t, err := template.ParseFiles("templates/chatgpt.page.tmpl")
		if err != nil {
			log.Println("Error to parse file chatgpt.page.tmpl", err)
		}
		t.Execute(w, nil)
	}
}
