package main

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.LoginTemplate)
	http.HandleFunc("/login/", authorization.Authorization)

	http.HandleFunc("/openai/", authorization.Authorization)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
