package completion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyCompletion struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	Suffix           string `json:"suffix"`
	MaxTokens        int    `json:"max_tokens"`
	Temperature      int    `json:"temperature"`
	TopP             int    `json:"top_p"`
	N                int    `json:"n"`
	Stream           bool   `json:"stream"`
	Logprobs         int    `json:"logprobs"`
	Echo             bool   `json:"echo"`
	Stop             string `json:"stop"`
	PresencePenalty  int    `json:"presence_penalty"`
	FrequencyPenalty int    `json:"frequency_penalty"`
	BestOf           int    `json:"best_of"`
	User             string `json:"user"`
}

type responseBodyCompletion struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func NewRequestBodyCompletion() RequestBodyCompletion {
	return RequestBodyCompletion{
		Model:            "text-davinci-003",
		Prompt:           "Say this is a test",
		Suffix:           "null",
		MaxTokens:        16,
		Temperature:      1,
		TopP:             1,
		N:                1,
		Stream:           false,
		Logprobs:         0,
		Echo:             false,
		Stop:             "null",
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		BestOf:           1,
		User:             "",
	}
}

var urlCompletion = "https://api.openai.com/v1/completions"

func CompletionOpenAI(apiKey string, req RequestBodyCompletion) (responseBodyCompletion, error) {

	reqBodyByte, _ := json.Marshal(req)

	r, err := http.NewRequest("POST", urlCompletion, bytes.NewBuffer(reqBodyByte))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	response := responseBodyCompletion{}

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil
}
