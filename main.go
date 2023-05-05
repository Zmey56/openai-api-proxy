package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
)

type responseBodyChat struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	apiKey := os.Getenv("API_KEY_OPENAI")

	remote, err := url.Parse("https://api.openai.com")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("r.URL", r.URL)
			r.Host = remote.Host
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			w.Header().Set("X-Ben", "Rad")
			buffer := bytes.NewBuffer([]byte{})
			writer := httptest.NewRecorder()
			p.ServeHTTP(writer, r)
			response := responseBodyChat{}
			derr := json.NewDecoder(writer.Body).Decode(&response)
			if derr != nil {
				panic(derr)
			}
			buffer.Write([]byte(response.Choices[0].Message.Content))
			w.Write(buffer.Bytes())
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", handler(proxy))

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
