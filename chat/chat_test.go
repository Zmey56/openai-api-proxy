package chat

import (
	"os"
	"testing"
)

func TestChatOpenAI(t *testing.T) {
	apiKey := os.Getenv("API_KEY_OPENAI")
	req := NewRequestBodyChart()
	resp, err := ChatOpenAI(apiKey, req)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if resp.ID == "" {
		t.Errorf("Expected ID to be populated, but it was empty")
	}
	if resp.Object != "chat.completion" {
		t.Errorf("Expected Object to be 'chat.completion', but it was %s", resp.Object)
	}
	if len(resp.Choices) != 1 {
		t.Errorf("Expected there to be 1 choice, but there were %d", len(resp.Choices))
	}
	if resp.Choices[0].Index != 0 {
		t.Errorf("Expected Index of first choice to be 0, but it was %d", resp.Choices[0].Index)
	}
	if resp.Choices[0].FinishReason != "stop" {
		t.Errorf("Expected FinishReason of first choice to be 'stop', but it was %s", resp.Choices[0].FinishReason)
	}
	if resp.Usage.PromptTokens == 0 {
		t.Errorf("Expected PromptTokens to be populated, but it was 0")
	}
	if resp.Usage.CompletionTokens == 0 {
		t.Errorf("Expected CompletionTokens to be populated, but it was 0")
	}
	if resp.Usage.TotalTokens == 0 {
		t.Errorf("Expected TotalTokens to be populated, but it was 0")
	}
}
