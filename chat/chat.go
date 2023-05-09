package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageChat struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

// RequestBodyChat remove LogitBias from request
type RequestBodyChat struct {
	Model            string        `json:"model"`
	Messages         []MessageChat `json:"messages"`
	Temperature      int           `json:"temperature"`
	TopP             int           `json:"top_p"`
	N                int           `json:"n"`
	Stream           bool          `json:"stream"`
	Stop             string        `json:"stop"`
	MaxTokens        int           `json:"max_tokens"`
	PresencePenalty  int           `json:"presence_penalty"`
	FrequencyPenalty int           `json:"frequency_penalty"`
	User             string        `json:"user"`
}

// responseBodyChat for request
type ResponseBodyChat struct {
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

// NewRequestBodyChart constructor for RequestBodyChat
func NewRequestBodyChart() RequestBodyChat {
	msg := MessageChat{
		Role:    "user",
		Content: "Hello!",
		Name:    "Alex",
	}
	return RequestBodyChat{
		Model:            "gpt-3.5-turbo-0301",
		Messages:         []MessageChat{msg},
		Temperature:      1,
		TopP:             1,
		N:                1,
		Stream:           false,
		Stop:             "null",
		MaxTokens:        4000,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		User:             "",
	}
}

var urlChat = "https://api.openai.com/v1/chat/completions"

func ChatOpenAI(apiKey string, req RequestBodyChat) (ResponseBodyChat, error) {

	reqBodyByte, _ := json.Marshal(req)

	r, err := http.NewRequest("POST", urlChat, bytes.NewBuffer(reqBodyByte))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	log.Println("Request", r)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	log.Println("Response", res)

	defer res.Body.Close()

	response := ResponseBodyChat{}

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil

}
