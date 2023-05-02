package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var urlChat = "https://api.openai.com/v1/chat/completions"

func main() {
	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Host)
	//if strings.HasPrefix(r.URL.String(), "/api") {
	//	//Check password
	//}
	//
	//url, err := url.Parse(fmt.Sprintf("http://%s/", urlChat))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//proxy := httputil.NewSingleHostReverseProxy(url)

	fmt.Println(r.URL.Host)
	apiKey := os.Getenv("API_KEY_OPENAI")

	if len(apiKey) < 1 {
		log.Println("You have problem with ApiKey")
	}

	data := map[string]interface{}{
		"model":  "gpt-3.5-turbo",
		"prompt": "Hello, world!",
	}

	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", urlChat, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println("resp", resp)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	//proxy.ServeHTTP(w, r)
}
