package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/audio"
	"github.com/Zmey56/openai-api-proxy/chat"
	"github.com/Zmey56/openai-api-proxy/completion"
	"github.com/Zmey56/openai-api-proxy/edit"
	"github.com/Zmey56/openai-api-proxy/embeddings"
	"github.com/Zmey56/openai-api-proxy/images"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type OpenAIRequest struct {
	Goal   string `json:"goal"`
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

func main() {

	apiKey := os.Getenv("API_KEY_OPENAI")

	remote, err := url.Parse("https://api.openai.com")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Host = remote.Host
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			w.Header().Set("X-Ben", "Rad")
			buffer := bytes.NewBuffer([]byte{})
			writer := httptest.NewRecorder()
			p.ServeHTTP(writer, r)
			response := make(map[string]interface{})
			derr := json.NewDecoder(writer.Body).Decode(&response)
			if derr != nil {
				panic(derr)
			}
			result := SwitchResponse(r.URL.String(), response)
			jsonByte, _ := json.Marshal(response)
			buffer.Write(jsonByte)
			w.Write(buffer.Bytes())
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", handler(proxy))

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}

func SwitchRequest(url string, param map[string]string) interface{} {
	urlArr := strings.Join(strings.Split(url, "/"), "")
	switch urlArr {
	case "v1completions":
		log.Println("v1/completions")
		reqComp := completion.NewRequestBodyCompletion()
		if param["model"] != "" {
			reqComp.Model = param["model"]
		}

		if param["prompt"] != "" {
			reqComp.Prompt = param["prompt"]
		}
		return reqComp
	case "v1chatcompletions":
		log.Println("v1/chat/completions")
		return chat.NewRequestBodyChart()
	case "v1edits":
		log.Println("v1/edits")
		return edit.NewRequestBodyEdit()
	case "v1imagesgenerations":
		log.Println("v1/images/generations")
		return images.NewRequestBodyImage()
	case "v1imagesedits":
		log.Println("v1/images/edits")
		return images.NewRequestBodyImageEdit()
	case "v1imagesvariations":
		log.Println("v1/images/variations")
		return images.NewRequestBodyImageVriation()
	case "v1embeddings":
		log.Println("v1/embeddings")
		return embeddings.NewRequestBodyEmbeddings()
	case "v1audiotranscriptions":
		log.Println("v1/audio/transcriptions")
		return audio.NewRequestBodyAudio()
	case "v1audiotranslations":
		log.Println("v1/audio/translations")
		return audio.NewRequestBodyAudioTranslation()
	default:
		log.Println("the method is not defined")
		return nil
	}
}

func SwitchResponse(url string, resp map[string]interface{}) interface{} {
	urlArr := strings.Join(strings.Split(url, "/"), "")
	//result := make(map[string]string)
	switch urlArr {
	case "v1completions":
		log.Println("v1/completions")
		return completion.ResponseBodyCompletion{}
	case "v1chatcompletions":
		log.Println("v1/chat/completions")
		return chat.ResponseBodyChat{}
	case "v1edits":
		log.Println("v1/edits")
		return edit.ResponseBodyEdit{}
	case "v1imagesgenerations":
		log.Println("v1/images/generations")
		return images.ResponseBodyImage{}
	case "v1imagesedits":
		log.Println("v1/images/edits")
		return images.ResponseBodyImage{}
	case "v1imagesvariations":
		log.Println("v1/images/variations")
		return images.ResponseBodyImage{}
	case "v1embeddings":
		log.Println("v1/embeddings")
		return embeddings.ResponseBodyEmbeddings{}
	case "v1audiotranscriptions":
		log.Println("v1/audio/transcriptions")
		return audio.ResponseBodyAudio{}
	case "v1audiotranslations":
		log.Println("v1/audio/translations")
		return audio.ResponseBodyAudio{}
	default:
		log.Println("the method is not defined")
		return nil
	}
}
