package edit

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestEditOpenAI(t *testing.T) {
	apiKey := os.Getenv("API_KEY_OPENAI")

	t.Run("Successful edit request", func(t *testing.T) {
		req := NewRequestBodyEdit()
		res, err := EditOpenAI(apiKey, req)
		if err != nil {
			t.Errorf("Error editing text: %v", err)
		}
		if res.Object != "edit" {
			t.Errorf("Expected 'object' to be 'edit', but got '%s'", res.Object)
		}
		if res.Created <= 0 {
			t.Errorf("Expected 'created' to be greater than 0, but got '%d'", res.Created)
		}
		if len(res.Choices) == 0 {
			t.Errorf("Expected 'choices' to have at least one element, but it is empty")
		}
		if res.Usage.PromptTokens <= 0 {
			t.Errorf("Expected 'prompt_tokens' to be greater than 0, but got '%d'", res.Usage.PromptTokens)
		}
		if res.Usage.CompletionTokens <= 0 {
			t.Errorf("Expected 'completion_tokens' to be greater than 0, but got '%d'", res.Usage.CompletionTokens)
		}
		if res.Usage.TotalTokens <= 0 {
			t.Errorf("Expected 'total_tokens' to be greater than 0, but got '%d'", res.Usage.TotalTokens)
		}
	})

	t.Run("Invalid API key", func(t *testing.T) {
		req := NewRequestBodyEdit()
		_, err := EditOpenAI("invalid_api_key", req)
		log.Println(err)
		if err == nil {
			t.Error("Expected an error, but got nil")
			return
		}
		if !strings.Contains(err.Error(), "Authentication failed") {
			t.Errorf("Expected 'Authentication failed' error message, but got '%v'", err)
		}
	})
}
