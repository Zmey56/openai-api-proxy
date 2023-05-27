package proxy_test

import (
	"encoding/json"
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/server/middlewares"
	"github.com/Zmey56/openai-api-proxy/server/proxy"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	t.Run("v1-chat-completions", func(t *testing.T) {
		openaiToken := "sg-foo"

		openAIMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v1/chat/completions" {
				t.Fatalf("unexpected path: %s", r.URL.Path)
			}

			if r.Header.Get("Authorization") != "Bearer "+openaiToken {
				t.Fatalf("unexpected token: %s", r.Header.Get("Authorization"))
			}

			if r.Header.Get("openai-api-proxy-user") != "test" {
				t.Fatalf("unexpected user: %s", r.Header.Get("openai-api-proxy-user"))
			}

			bodyJSON := make(map[string]interface{})
			err := json.NewDecoder(r.Body).Decode(&bodyJSON)
			if err != nil {
				t.Fatal(err)
			}
			if bodyJSON["user"] != "test" {
				t.Fatalf("unexpected user: %s", bodyJSON["user"])
			}
			if bodyJSON["model"] != "gpt-3.5-turbo" {
				t.Fatalf("unexpected model: %s", bodyJSON["model"])
			}
			messages := bodyJSON["messages"].([]interface{})
			if len(messages) != 1 {
				t.Fatalf("unexpected messages length: %d", len(messages))
			}
			message := messages[0].(map[string]interface{})
			if message["role"] != "user" {
				t.Fatalf("unexpected message role: %s", message["role"])
			}
			if message["content"] != "Hello!" {
				t.Fatalf("unexpected message content: %s", message["content"])
			}

			w.Header().Set("Openai-Model", "gpt-3.5-turbo-0301")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			_, err = w.Write([]byte(`{"id":"chatcmpl-7El6pSCt0PoA1a9JWuIFzDNwwXn3B","object":"chat.completion","created":1683752035,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":10,"completion_tokens":10,"total_tokens":20},"choices":[{"message":{"role":"assistant","content":"Hello there! How can I assist you today?"},"finish_reason":"stop","index":0}]}`))
			if err != nil {
				t.Fatal(err)
			}
		}))
		defer openAIMock.Close()

		openAIURL := openAIMock.URL

		p, err := proxy.NewProxy(proxy.Configuration{
			OpenaiAddress: openAIURL,
			OpenaiToken:   openaiToken,
		})

		if err != nil {
			t.Fatal(err)
		}

		// TODO: fix test!!!
		proxyServer := httptest.NewServer(middlewares.RemovePathPrefixMiddleware(
			middlewares.AuthorizationMiddleware(p, authorization.StaticService{}, nil),
			"/openai/",
		))
		defer proxyServer.Close()

		urlProxyServer, err := url.Parse(proxyServer.URL)
		if err != nil {
			t.Fatal(err)
		}
		urlProxyServer.Path = "/openai/v1/chat/completions"
		request, err := http.NewRequest(http.MethodPost, urlProxyServer.String(), strings.NewReader(`{"model": "gpt-3.5-turbo","messages":[{"role": "user", "content": "Hello!"}]}`))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "test")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status code: %d", response.StatusCode)
		}

		bodyJSON := make(map[string]interface{})
		err = json.NewDecoder(response.Body).Decode(&bodyJSON)
		if err != nil {
			t.Fatal(err)
		}
		if bodyJSON["id"] != "chatcmpl-7El6pSCt0PoA1a9JWuIFzDNwwXn3B" {
			t.Fatalf("unexpected id: %s", bodyJSON["id"])
		}
		if bodyJSON["object"] != "chat.completion" {
			t.Fatalf("unexpected object: %s", bodyJSON["object"])
		}
		if bodyJSON["created"] != 1683752035.0 {
			t.Fatalf("unexpected created: %f", bodyJSON["created"])
		}
		if bodyJSON["model"] != "gpt-3.5-turbo-0301" {
			t.Fatalf("unexpected model: %s", bodyJSON["model"])
		}
		if bodyJSON["usage"].(map[string]interface{})["prompt_tokens"] != 10.0 {
			t.Fatalf("unexpected prompt_tokens: %f", bodyJSON["usage"].(map[string]interface{})["prompt_tokens"])
		}
		if bodyJSON["usage"].(map[string]interface{})["completion_tokens"] != 10.0 {
			t.Fatalf("unexpected completion_tokens: %f", bodyJSON["usage"].(map[string]interface{})["completion_tokens"])
		}
		if bodyJSON["usage"].(map[string]interface{})["total_tokens"] != 20.0 {
			t.Fatalf("unexpected total_tokens: %f", bodyJSON["usage"].(map[string]interface{})["total_tokens"])
		}
		choices := bodyJSON["choices"].([]interface{})
		if len(choices) != 1 {
			t.Fatalf("unexpected choices length: %d", len(choices))
		}
		choice := choices[0].(map[string]interface{})
		if choice["finish_reason"] != "stop" {
			t.Fatalf("unexpected finish_reason: %s", choice["finish_reason"])
		}
		if choice["index"] != 0.0 {
			t.Fatalf("unexpected index: %f", choice["index"])
		}
		message := choice["message"].(map[string]interface{})
		if message["role"] != "assistant" {
			t.Fatalf("unexpected message role: %s", message["role"])
		}
		if message["content"] != "Hello there! How can I assist you today?" {
			t.Fatalf("unexpected message content: %s", message["content"])
		}
	})
}
