package models

import (
	"os"
	"testing"
)

func TestModelsOpenAI(t *testing.T) {
	apiKey := os.Getenv("API_KEY_OPENAI")

	response, err := ModelsOpenAI(apiKey)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(response.Data) == 0 {
		t.Errorf("Expected at least one model, but got non")
	}

	//Test finding an existing model
	modelName := "davinci"
	response.FindModel(modelName)

	// Test finding a non-existing model
	modelName = "nonexistent"
	response.FindModel(modelName)
}

func TestRetriveModelOpenAI(t *testing.T) {
	// Replace with your own API key and model name
	apiKey := os.Getenv("API_KEY_OPENAI")
	model := "davinci"

	response, err := RetriveModelOpenAI(apiKey, model)

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	// Check if the response ID is not empty
	if response.Id == "" {
		t.Errorf("Response ID should not be empty")
	}

	// Check if the response object is "model"
	if response.Object != "model" {
		t.Errorf("Response object should be 'model'")
	}

	// Check if the owned_by field is not empty
	if response.OwnedBy == "" {
		t.Errorf("owned_by field should not be empty")
	}

	// Check if there is at least one permission in the response
	if len(response.Permission) < 1 {
		t.Errorf("Response should contain at least one permission")
	}
}
