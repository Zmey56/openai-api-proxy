package audio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type RequestBodyAudio struct {
	File           string  `json:"file"`
	Model          string  `json:"model"`
	Prompt         string  `json:"prompt"`
	ResponseFormat string  `json:"response_format"`
	Temperature    float64 `json:"temperature"`
	Language       string  `json:"language"`
}

type RequestBodyAudioTranslation struct {
	File           string  `json:"file"`
	Model          string  `json:"model"`
	Prompt         string  `json:"prompt"`
	ResponseFormat string  `json:"response_format"`
	Temperature    float64 `json:"temperature"`
}

type ResponseBodyAudio struct {
	Text string `json:"text"`
}

func NewRequestBodyAudio() RequestBodyAudio {
	return RequestBodyAudio{
		File:           "audio.mp3",
		Model:          "whisper-1",
		Prompt:         "Audio",
		ResponseFormat: "json",
		Temperature:    0,
		Language:       "en",
	}
}

func NewRequestBodyAudioTranslation() RequestBodyAudioTranslation {
	return RequestBodyAudioTranslation{
		File:           "audio.mp3",
		Model:          "whisper-1",
		Prompt:         "Audio",
		ResponseFormat: "json",
		Temperature:    0,
	}
}

var urlAudio = "https://api.openai.com/v1/audio/transcriptions"

func TranscriptionsOpenAI(apiKey string) (ResponseBodyAudio, error) {
	resBA := ResponseBodyAudio{}
	//
	//requestBody := RequestBodyAudio{}

	nameFile := "audio.m4a"

	path, _ := os.Getwd()
	pathFile := fmt.Sprintf("%s/audio/%s", path, nameFile)
	log.Println(pathFile)

	file, err := os.Open(pathFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multipartWriter.CreateFormFile("file", pathFile)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}

	modelFieldWriter, err := multipartWriter.CreateFormField("model")
	if err != nil {
		panic(err)
	}
	modelFieldWriter.Write([]byte("whisper-1"))

	multipartWriter.Close()

	req, err := http.NewRequest("POST", urlAudio, &requestBody)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))

	derr := json.NewDecoder(response.Body).Decode(&resBA)
	if derr != nil {
		return resBA, err
	}

	if response.StatusCode != http.StatusOK {
		return resBA, err
	}

	log.Println(response)

	return resBA, nil

}
