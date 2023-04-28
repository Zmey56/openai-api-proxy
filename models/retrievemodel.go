package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type responseBodyModel struct {
	Id         string `json:"id"`
	Object     string `json:"object"`
	Created    int    `json:"created"`
	OwnedBy    string `json:"owned_by"`
	Permission []struct {
		Id                 string      `json:"id"`
		Object             string      `json:"object"`
		Created            int         `json:"created"`
		AllowCreateEngine  bool        `json:"allow_create_engine"`
		AllowSampling      bool        `json:"allow_sampling"`
		AllowLogprobs      bool        `json:"allow_logprobs"`
		AllowSearchIndices bool        `json:"allow_search_indices"`
		AllowView          bool        `json:"allow_view"`
		AllowFineTuning    bool        `json:"allow_fine_tuning"`
		Organization       string      `json:"organization"`
		Group              interface{} `json:"group"`
		IsBlocking         bool        `json:"is_blocking"`
	} `json:"permission"`
	Root   string      `json:"root"`
	Parent interface{} `json:"parent"`
}

func RetriveModelOpenAI(apiKey, model string) (responseBodyModel, error) {
	urlModels := fmt.Sprintf("%s/%s", urlModels, model)
	log.Println(urlModels)

	r, err := http.NewRequest("GET", urlModels, nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	response := responseBodyModel{}

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode == http.StatusNotFound {
		fmt.Println("Model not found")
		return response, nil
	} else if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil

}
