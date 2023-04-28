package files

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var urlFile = "https://api.openai.com/v1/files"

type responseBodyChat struct {
	Data []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"data"`
	Object string `json:"object"`
}

func FilesOpenAI(apiKey string) (responseBodyChat, error) {
	response := responseBodyChat{}

	r, err := http.NewRequest("POST", urlFile, nil)
	//r.Header.Add("Content-Type", "multipart/form-data")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()
	log.Println(res)

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		return response, derr
	}

	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil

}
