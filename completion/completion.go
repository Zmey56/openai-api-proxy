package completion

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RequestBodyCompletion struct {
	Model       string      `json:"model"`
	Prompt      string      `json:"prompt"`
	MaxTokens   int         `json:"max_tokens"`
	Temperature int         `json:"temperature"`
	TopP        int         `json:"top_p"`
	N           int         `json:"n"`
	Stream      bool        `json:"stream"`
	Logprobs    interface{} `json:"logprobs"`
	Stop        string      `json:"stop"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
	Usage struct {
		Prompt_tokens     int `json:"prompt_tokens"`
		Completion_tokens int `json:"completion_tokens"`
		Total_tokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type OpenAIRequest struct {
	Prompt           string  `json:"prompt"`
	Model            string  `json:"model"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

var urlCompletion = "https://api.openai.com/v1/completions"

func CompletionOpenAI(apiKey string) {
	requestTest := OpenAIRequest{}
	requestTest.Prompt = "test"
	requestTest.Model = "text-davinci-002"
	requestTest.MaxTokens = 100
	requestTest.Temperature = 0.5

	requestBodyTest, _ := json.Marshal(requestTest)

	requestBodyBytes, _ := json.Marshal(requestBodyTest)
	request, err := http.NewRequest("POST", urlCompletion, strings.NewReader(string(requestBodyBytes)))
	if err != nil {
		//http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		//http.Error(w, "Failed to send request to OpenAI API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)

	// Response processing
	var openAIResponse OpenAIResponse
	err = json.NewDecoder(response.Body).Decode(&openAIResponse)
	if err != nil {
		//http.Error(w, "Failed to parse response from OpenAI API", http.StatusInternalServerError)
		return
	}

	log.Println(openAIResponse)

	//requestBody := RequestBodyCompletion{}
	//requestBody.Prompt = "Say this is a test"
	//requestBody.Model = "text-davinci-002"
	//requestBody.MaxTokens = 7
	//requestBody.Temperature = 0
	//requestBody.TopP = 1
	//requestBody.N = 1
	//requestBody.Stream = false
	//requestBody.Logprobs = "null"
	//requestBody.Stop = "\n"
	//
	//requestBodyBytes, _ := json.Marshal(requestBody)
	//request, err := http.NewRequest("POST", urlCompletion, strings.NewReader(string(requestBodyBytes)))
	//if err != nil {
	//	log.Println("Failed to create request", err)
	//	return
	//}
	//request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	//
	//client := &http.Client{}
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	log.Println("Failed to send request to OpenAI API", err)
	//	return
	//}
	//log.Println(response.StatusCode)
	//defer response.Body.Close()
	//
	//// Response processing
	//var openAIResponse OpenAIResponse
	//err = json.NewDecoder(response.Body).Decode(&openAIResponse)
	//if err != nil {
	//	log.Println("Failed to parse response from OpenAI API", err)
	//	return
	//}
	//
	////How many user used tokens
	//numTokens := openAIResponse.Usage.Prompt_tokens
	//log.Println(numTokens)
	//
	//log.Println(openAIResponse)

}
