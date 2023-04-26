package embeddings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyEmbeddings struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user"`
}

type responseBodyEmbeddings struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Index     int       `json:"index"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

var urlEmbendding = "https://api.openai.com/v1/embeddings"

func EmbenddingOpenAI(apiKey string) (responseBodyEmbeddings, error) {
	response := responseBodyEmbeddings{}

	requestBody := RequestBodyEmbeddings{}
	//
	requestBody.Model = "text-embedding-ada-002"
	requestBody.Input = "The food was delicious and the waiter..."

	reqBodyByte, _ := json.Marshal(requestBody)

	r, err := http.NewRequest("POST", urlEmbendding, bytes.NewBuffer(reqBodyByte))
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

	log.Println(response)

	return response, nil
}
