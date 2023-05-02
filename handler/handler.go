package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var openAihost = "https://api.openai.com/v1/chat/completions"

func HandlerProxy(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Host)

	apiKey := os.Getenv("API_KEY_OPENAI")

	if len(apiKey) < 1 {
		log.Println("You have problem with ApiKey")
	}

	tokenValue := r.Header.Get("Authorization")
	if !checkAuthorization(tokenValue) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	url, err := url.Parse(openAihost)
	if err != nil {
		log.Println("Error with url", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	log.Println(r.URL.Host)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	proxy.ServeHTTP(w, r)
}

func checkAuthorization(t string) bool {
	return true
}
