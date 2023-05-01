package edit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBodyEdit struct {
	Model       string `json:"model"`
	Input       string `json:"input"`
	Instruction string `json:"instruction"`
	N           int    `json:"n"`
	Temperature int    `json:"temperature"`
	TopP        int    `json:"top_p"`
}

type responseBodyEdit struct {
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Text  string `json:"text"`
		Index int    `json:"index"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func NewRequestBodyEdit() RequestBodyEdit {
	return RequestBodyEdit{
		Model:       "text-davinci-edit-001",
		Input:       "What day of the wek is it?",
		Instruction: "Fix the spelling mistakes",
		N:           1,
		Temperature: 1,
		TopP:        1,
	}
}

var urlEdit = "https://api.openai.com/v1/edits"

func EditOpenAI(apiKey string, req RequestBodyEdit) (responseBodyEdit, error) {
	response := responseBodyEdit{}

	reqBodyByte, _ := json.Marshal(req)

	r, err := http.NewRequest("POST", urlEdit, bytes.NewBuffer(reqBodyByte))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		return response, err
	}

	if res.StatusCode != http.StatusOK {
		return response, err
	}

	return response, nil
}
