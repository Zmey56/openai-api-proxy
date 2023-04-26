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
	"path/filepath"
)

type RequestBodyAudio struct {
	File           string `json:"file"`
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	ResponseFormat string `json:"response_format"`
	Temperature    int    `json:"temperature"`
	Language       string `json:"language"`
}

type responseBodyAudio struct {
	Text string `json:"text"`
}

var urlAudio = "https://api.openai.com/v1/audio/transcriptions"

func AudioOpenAI(apiKey string) (responseBodyAudio, error) {
	response := responseBodyAudio{}

	requestBody := RequestBodyAudio{}

	nameFile := "audio.m4a"

	path, _ := os.Getwd()
	pathFile := fmt.Sprintf("%s/audion/%s", path, nameFile)
	log.Println(pathFile)
	file, _ := os.Open(pathFile)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	log.Println("body", body)
	log.Println("writer", writer)
	log.Println("part", part)
	log.Println("file", file)

	//
	requestBody.Model = "whisper-1" //only one
	requestBody.File = pathFile

	reqBodyByte, _ := json.Marshal(requestBody)

	r, err := http.NewRequest("POST", urlAudio, bytes.NewBuffer(reqBodyByte))
	r.Header.Add("Content-Type", "multipart/form-data")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return response, err
	}

	defer res.Body.Close()

	log.Println("RESPONSE", res)

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		log.Println(err)
		return response, err
	}

	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return response, err
	}

	log.Println(response)

	return response, nil

}
