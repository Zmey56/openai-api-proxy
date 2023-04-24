package chat

import (
	"fmt"
)

type requestBody struct {
	model            string              `json:"model"`
	messages         []map[string]string `json:"messages"`
	temperature      int                 `json:"temperature"`
	topP             int                 `json:"top_p"`
	n                int                 `json:"n"`
	stream           bool                `json:"stream"`
	stop             string              `json:"stop"`
	maxTokens        int                 `json:"max_tokens"`
	presencePenalty  int                 `json:"presence_penalty"`
	frequencyPenalty int                 `json:"frequency_penalty"`
	logitBias        map[string]string   `json:"logit_bias"`
	user             string              `json:"user"`
}

type responseBody struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

var urlChat = "https://api.openai.com/v1/chat/completions"

func ChatOpenAI(key string) {
	// authorization on OpenAI
	authHeader := fmt.Sprintf("Bearer %s", key)

	//req, err := http.NewRequest("POST", urlChat, )

}
