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

type RequestBodyChat struct {
	Model            string         `json:"model"`
	Messages         []MessageChat  `json:"messages"`
	Temperature      int            `json:"temperature"`
	TopP             int            `json:"top_p"`
	N                int            `json:"n"`
	Stream           bool           `json:"stream"`
	Stop             string         `json:"stop"`
	MaxTokens        int            `json:"max_tokens"`
	PresencePenalty  int            `json:"presence_penalty"`
	FrequencyPenalty int            `json:"frequency_penalty"`
	LogitBias        map[string]int `json:"logit_bias"`
	User             string         `json:"user"`
}

type SmallRequestBodyChat struct {
	Model    string        `json:"model"`
	Messages []MessageChat `json:"messages"`
}

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

var urlChat = "https://api.openai.com/v1/chat/completions"

func ChatOpenAI(apiKey string) (responseBodyChat, error) {

	// gpt-4, gpt-4-0314, gpt-4-32k, gpt-4-32k-0314, gpt-3.5-turbo, gpt-3.5-turbo-0301
	reqBody := &SmallRequestBodyChat{
		Model: "gpt-3.5-turbo",
		Messages: []MessageChat{{
			Role:    "user",
			Content: "Hello!",
			Name:    "TestUser", //can use name of user
		}},
	}

	reqBodyByte, _ := json.Marshal(reqBody)

	r, err := http.NewRequest("POST", urlChat, bytes.NewBuffer(reqBodyByte))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	response := responseBodyChat{}

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil

}
