package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type MessageChat struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type RequestBodyChat struct {
	Model            string            `json:"model"`
	Messages         []MessageChat     `json:"messages"`
	Temperature      int               `json:"temperature"`
	TopP             int               `json:"top_p"`
	N                int               `json:"n"`
	Stream           bool              `json:"stream"`
	Stop             string            `json:"stop"`
	MaxTokens        int               `json:"max_tokens"`
	PresencePenalty  int               `json:"presence_penalty"`
	FrequencyPenalty int               `json:"frequency_penalty"`
	LogitBias        map[string]string `json:"logit_bias"`
	User             string            `json:"user"`
}

type responseBody struct {
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

func ChatOpenAI(apiKey string) {

	// gpt-4, gpt-4-0314, gpt-4-32k, gpt-4-32k-0314, gpt-3.5-turbo, gpt-3.5-turbo-0301
	reqBody := RequestBodyChat{
		Model: "gpt-3.5-turbo",
		Messages: []MessageChat{{
			Role:    "user",
			Content: "Hello!",
			Name:    "Test",
		}},
	}

	log.Println("reqBody", reqBody)

	requestBodyBytes, _ := json.Marshal(reqBody)
	request, err := http.NewRequest("POST", urlChat, strings.NewReader(string(requestBodyBytes)))
	if err != nil {
		log.Println("Failed to create request", err)
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)

	// read the response body
	var respData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		fmt.Println("Error decoding response data:", err)
		return
	}
	// print the response data
	fmt.Println(respData)

}
