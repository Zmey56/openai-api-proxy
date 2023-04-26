package images

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBodyImage struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type SmallRequestBodyImage struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type responseBodyImage struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

var urlImage = "https://api.openai.com/v1/images/generations"

func ImageOpenAI(apiKey string) (responseBodyImage, error) {
	response := responseBodyImage{}

	requestBody := SmallRequestBodyImage{}
	//
	requestBody.Prompt = "Zmey56, Golang, Python Developer, Data Analytics, Quantitative Analyst, Data Science, BI Analyst, Blockchain, Cryptocurrency"
	requestBody.N = 2
	requestBody.Size = "1024x1024"

	reqBodyByte, _ := json.Marshal(requestBody)

	r, err := http.NewRequest("POST", urlImage, bytes.NewBuffer(reqBodyByte))
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
