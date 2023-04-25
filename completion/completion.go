package completion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyCompletion struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix"`
	MaxTokens        int            `json:"max_tokens"`
	Temperature      int            `json:"temperature"`
	TopP             int            `json:"top_p"`
	N                int            `json:"n"`
	Stream           bool           `json:"stream"`
	Logprobs         interface{}    `json:"logprobs"`
	Echo             bool           `json:"echo"`
	Stop             string         `json:"stop"`
	PresencePenalty  int            `json:"presence_penalty"`
	FrequencyPenalty int            `json:"frequency_penalty"`
	BestOf           int            `json:"best_of"`
	LogitBias        map[string]int `json:"logit_bias"`
	User             string         `json:"user"`
}

type SmallRequestBodyCompletion struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
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

var urlCompletion = "https://api.openai.com/v1/completions"

func CompletionOpenAI(apiKey string) (responseBodyCompletion, error) {
	requestBody := SmallRequestBodyCompletion{}
	requestBody.Prompt = "test"
	// text-davinci-003, text-davinci-002, text-curie-001, text-babbage-001, text-ada-001
	requestBody.Model = "text-davinci-002"

	reqBodyByte, _ := json.Marshal(requestBody)

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
