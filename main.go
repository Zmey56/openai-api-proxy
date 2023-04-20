package main

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/handlers"
	"log"
	"net/http"
)

var tokens = make(map[string]authorization.Token)

func main() {

	http.HandleFunc("/", handlers.LoginTemplate)
	http.HandleFunc("/login/", authorization.Authorization(tokens))
	http.HandleFunc("/openai", handlers.ChatGPT(tokens))
	//http.HandleFunc("/openai", handlers.ChatGPT(tokens))
	//http.HandleFunc("/openai/", authorization.Authorization(tokens))

	log.Fatal(http.ListenAndServe(":8081", nil))

}
