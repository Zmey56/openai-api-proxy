package main

import (
	"flag"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/chat"
	"github.com/Zmey56/openai-api-proxy/completion"
	"log"
	"os"
	"strings"
)

var (
	defaultGoal  = "chat"
	allowedGoals = []string{"list", "completion", "chat", "edits",
		"images", "embeddings", "audio", "files", "finetune", "moderations"}
	selectedGoal string
)

func init() {
	flag.StringVar(&selectedGoal, "goals", defaultGoal, fmt.Sprintf("select a model from: %v ", strings.Join(allowedGoals, ", ")))
	flag.StringVar(&selectedGoal, "g", defaultGoal, fmt.Sprintf("'select a model from: %v'", strings.Join(allowedGoals, ", ")))
}

func main() {

	flag.Parse()

	apiKey := os.Getenv("API_KEY_OPENAI")
	//if len(apiKey) < 1 {
	//	log.Fatal("Not api key")
	//}

	log.Println(selectedGoal)

	switch selectedGoal {
	case "list":
		fmt.Println("LIST")
	case "completion":
		fmt.Println("completion")
		completion.CompletionOpenAI(apiKey)
	case "chat":
		fmt.Println("chat")
		chat.ChatOpenAI(apiKey)
	case "edits":
		fmt.Println("edits")
	case "images":
		fmt.Println("images")
	case "embeddings":
		fmt.Println("embeddings")
	case "audio":
		fmt.Println("audio")
	case "files":
		fmt.Println("files")
	case "finetune":
		fmt.Println("finetune")
	case "moderations":
		fmt.Println("moderations")
	default:
		fmt.Println("Invalid model selected")

	}
	//
	//// List models
	//urlListModels := "https://api.openai.com/v1/models"
	//
	//// Create completion
	//urlComp := "https://api.openai.com/v1/completions"
	//
	//// Chat
	//urlChat := "https://api.openai.com/v1/chat/completions"
	//
	//// Edits
	//urlEdits := "https://api.openai.com/v1/edits"
	//
	//// Images
	//urlImag := "https://api.openai.com/v1/images/generations"
	//
	//// Embeddings
	//urlEmb := "https://api.openai.com/v1/embeddings"
	//
	//// Audio
	//urlAudio := "https://api.openai.com/v1/audio/transcriptions"
	//
	//// Files
	//urlFiles := "https://api.openai.com/v1/files"
	//
	//// Fine-tune
	//urlFineTune := "https://api.openai.com/v1/fine-tunes"
	//
	//// Moderations
	//urlModer := "https://api.openai.com/v1/moderations"

}
