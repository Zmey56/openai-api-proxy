package main

import (
	"flag"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/audio"
	"github.com/Zmey56/openai-api-proxy/chat"
	"github.com/Zmey56/openai-api-proxy/completion"
	"github.com/Zmey56/openai-api-proxy/edit"
	"github.com/Zmey56/openai-api-proxy/embeddings"
	"github.com/Zmey56/openai-api-proxy/files"
	"github.com/Zmey56/openai-api-proxy/images"
	"github.com/Zmey56/openai-api-proxy/models"
	"log"
	"os"
	"strings"
)

var (
	defaultGoal  = "models"
	allowedGoals = []string{"models", "list", "completion", "chat", "edits",
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
	if len(apiKey) < 1 {
		log.Fatal("Not api key")
	}

	args := os.Args
	log.Println("args", args[1])

	switch selectedGoal {
	case "models":
		fmt.Println("models")
		if len(args) > 2 {
			log.Println("FIND")
			model, _ := models.RetriveModelOpenAI(apiKey, strings.TrimSpace(args[2]))
			log.Println(model)
		} else {
			models, _ := models.ModelsOpenAI(apiKey)
			models.FindModel("text-ada-001")
			models.FindModel("text-zmey-001")
		}
	case "completion":
		fmt.Println("completion")
		completion.CompletionOpenAI(apiKey)
	case "chat":
		fmt.Println("chat")
		chat.ChatOpenAI(apiKey)
	case "edits":
		fmt.Println("edits")
		edit.EditOpenAI(apiKey)
	case "images":
		fmt.Println("images")
		images.ImageOpenAI(apiKey)
	case "embeddings":
		fmt.Println("embeddings")
		embeddings.EmbenddingOpenAI(apiKey)
	case "audio":
		fmt.Println("audio")
		audio.TranscriptionsOpenAI(apiKey)
	case "files":
		fmt.Println("files")
		//files.FilesOpenAI(apiKey)
		files.UploadOpenAI(apiKey)
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
