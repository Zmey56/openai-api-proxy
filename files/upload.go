package files

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

type RequestBodyUpload struct {
	File    string `json:"file"`
	Purpose string `json:"purpose"`
}

type responseBodyUpload struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

var urlUpload = "https://api.openai.com/v1/files"

func UploadOpenAI(apiKey string) (responseBodyUpload, error) {
	resU := responseBodyUpload{}

	nameFile := "test.jsonl"

	path, _ := os.Getwd()
	pathFile := fmt.Sprintf("%s/files/%s", path, nameFile)
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

	modelFieldWriter, err := multipartWriter.CreateFormField("purpose")
	if err != nil {
		panic(err)
	}
	modelFieldWriter.Write([]byte("fine-tune"))

	multipartWriter.Close()

	req, err := http.NewRequest("POST", urlUpload, &requestBody)
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

	derr := json.NewDecoder(response.Body).Decode(&resU)
	if derr != nil {
		return resU, err
	}

	if response.StatusCode != http.StatusOK {
		return resU, err
	}

	log.Println(response)

	return resU, nil
}
