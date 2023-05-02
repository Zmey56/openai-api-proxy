package completion

import (
	"os"
	"testing"
)

func TestCompletionOpenAI(t *testing.T) {
	// Set up a fake API key and request body
	apiKey := os.Getenv("API_KEY_OPENAI")
	reqBody := NewRequestBodyCompletion()

	// Call the function being tested
	res, err := CompletionOpenAI(apiKey, reqBody)

	// Check that the function didn't return an error
	if err != nil {
		t.Errorf("CompletionOpenAI returned an error: %v", err)
	}

	// Check that the response object has the expected fields
	if res.Id == "" {
		t.Errorf("Response object doesn't have expected Id field")
	}
	if res.Object == "" {
		t.Errorf("Response object doesn't have expected Object field")
	}
	if res.Created == 0 {
		t.Errorf("Response object doesn't have expected Created field")
	}
	if res.Model != reqBody.Model {
		t.Errorf("Response object has unexpected Model field value")
	}
	if len(res.Choices) != 1 {
		t.Errorf("Response object doesn't have expected number of Choices")
	}
	if res.Choices[0].Text == "" {
		t.Errorf("Response object has empty Text field in Choices")
	}
	if res.Choices[0].Index != 0 {
		t.Errorf("Response object has unexpected Index field value")
	}
	if res.Choices[0].FinishReason == "" {
		t.Errorf("Response object has empty FinishReason field in Choices")
	}
	if res.Usage.PromptTokens == 0 {
		t.Errorf("Response object has unexpected PromptTokens field value")
	}
	if res.Usage.CompletionTokens == 0 {
		t.Errorf("Response object has unexpected CompletionTokens field value")
	}
	if res.Usage.TotalTokens == 0 {
		t.Errorf("Response object has unexpected TotalTokens field value")
	}
}
