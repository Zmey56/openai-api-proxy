package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type responseBodyModels struct {
	Object string `json:"object"`
	Data   []struct {
		ID         string `json:"id"`
		Object     string `json:"object"`
		Created    int64  `json:"created"`
		OwnedBy    string `json:"owned_by"`
		Permission []struct {
			ID                 string `json:"id"`
			Object             string `json:"object"`
			Created            int    `json:"created"`
			AllowCreateEngine  bool   `json:"allow_create_engine"`
			AllowSampling      bool   `json:"allow_sampling"`
			AllowLogprobs      bool   `json:"allow_logprobs"`
			AllowSearchIndices bool   `json:"allow_search_indices"`
			AllowView          bool   `json:"allow_view"`
			AllowFineTuning    bool   `json:"allow_fine_tuning"`
			Organization       string `json:"organization"`
			Group              any    `json:"group"`
			IsBlocking         bool   `json:"is_blocking"`
		} `json:"permission"`
		Root   string `json:"root"`
		Parent any    `json:"parent"`
	} `json:"data"`
}

var urlModels = "https://api.openai.com/v1/models"

func ModelsOpenAI(apiKey string) (responseBodyModels, error) {

	r, err := http.NewRequest("GET", urlModels, nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	response := responseBodyModels{}

	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}

	return response, nil

}

func (m *responseBodyModels) FindModel(name string) {
	findMod := false
	for _, model := range m.Data {
		if model.ID == name {
			t := time.Unix(model.Created, 0)
			formattedDate := t.Format("2006 January, 02")
			fmt.Println("Model:", model.ID, "Created", formattedDate, "Owned", model.OwnedBy)
			fmt.Println("Permission:")
			for _, al := range model.Permission {
				fmt.Println("- allow create engine:", al.AllowCreateEngine)
				fmt.Println("- allow sampling:", al.AllowSampling)
				fmt.Println("- allow logprobs:", al.AllowLogprobs)
				fmt.Println("- allow search indices:", al.AllowSearchIndices)
				fmt.Println("- allow view:", al.AllowView)
			}
			findMod = true
			break
		}
	}
	if !findMod {
		fmt.Println("Model not found")
	}
}
