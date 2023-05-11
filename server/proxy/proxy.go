package proxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/log"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type openAIResponse struct {
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
}

func NewProxy(conf Configuration) (*Proxy, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	openaiURL, err := url.Parse(conf.OpenaiAddress)
	if err != nil {
		return nil, err
	}

	p := &Proxy{
		id64Generator: newID64Generator(),
		openaiURL:     openaiURL,
		conf:          conf,
	}

	p.proxy = &httputil.ReverseProxy{
		Rewrite:        p.rewrite,
		ModifyResponse: p.modifyResponse,
		ErrorHandler:   p.errorHandler,
	}

	return p, nil
}

type Proxy struct {
	id64Generator *id64Generator

	openaiURL *url.URL

	conf Configuration

	proxy *httputil.ReverseProxy
}

func (s *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func (s *Proxy) rewrite(request *httputil.ProxyRequest) {
	requestID := s.id64Generator.New().String()

	if log.IsTrace() {
		if requestText, err := httputil.DumpRequest(request.In, true); err == nil {
			log.Trace.Printf("request in %s: \n%s", requestID, requestText)
		} else {
			log.Trace.Printf("failed to dump request in %s: %s", requestID, err.Error())
		}
	}

	request.SetURL(s.openaiURL)
	request.SetXForwarded()
	request.Out.Header.Set("opeanai-proxy-request-id", requestID)
	request.Out.Header.Set("Authorization", "Bearer "+s.conf.OpenaiToken)

	if request.In.Header.Get("Content-Type") != "application/json" {
		return
	}

	if request.In.Body == nil || request.In.ContentLength == 0 {
		return
	}

	//In case if body has JSON, we will inject user in the body
	user := request.In.Header.Get("openai-api-proxy-user")

	body, err := io.ReadAll(request.In.Body)
	if err != nil {
		log.Warning.Printf("failed to read the body of the incoming request - %s - %s", request.In.URL.Path, err.Error())
		return
	}

	result, err := injectUser(bytes.NewBuffer(body), user)
	if err != nil {
		log.Warning.Printf("failed to inject user into the body of the incoming request - %s - %s", request.In.URL.Path, err.Error())
		return
	}

	request.Out.ContentLength = int64(result.Len())
	request.Out.Body = io.NopCloser(result)

	if log.IsTrace() {
		if requestText, err := httputil.DumpRequest(request.Out, true); err == nil {
			log.Trace.Printf("request out %s: \n%s", requestID, requestText)
		} else {
			log.Trace.Printf("failed to dump request out %s: %s", requestID, err.Error())
		}
	}
}

func (s *Proxy) modifyResponse(response *http.Response) error {
	requestID := response.Request.Header.Get("opeanai-proxy-request-id")

	if log.IsTrace() {
		if requestText, err := httputil.DumpResponse(response, true); err == nil {
			log.Trace.Printf("response %s: \n%s", requestID, requestText)
		} else {
			log.Trace.Printf("failed to dump request %s: %s", requestID, err.Error())
		}
	}

	user := response.Request.Header.Get("openai-api-proxy-user")

	if response.Header.Get("Content-Type") != "application/json" {
		log.Info.Printf("usage - %s - %s - path: %s", requestID, user, response.Request.URL.Path)
		return nil
	}

	openAIModel := response.Header.Get("Openai-Model")

	if response.StatusCode != http.StatusOK {
		return nil
	}

	// just trying to decode the response body and get some usage of the tokens
	body := bytes.Buffer{}
	bodyReader := io.TeeReader(response.Body, &body)
	responseObj := openAIResponse{}
	if err := json.NewDecoder(bodyReader).Decode(&responseObj); err != nil {
		return errors.Join(err, fmt.Errorf("failed to decode response"))
	}

	logString := strings.Builder{}
	logString.WriteString("usage - ")
	logString.WriteString(requestID)
	logString.WriteString(" - ")
	logString.WriteString(response.Request.URL.Path)
	logString.WriteString(" - ")
	logString.WriteString(user)
	logString.WriteString(" - ")
	logString.WriteString(openAIModel)
	logString.WriteString(" -")
	if responseObj.Usage.TotalTokens > 0 {
		logString.WriteString(
			fmt.Sprintf(
				" prompt=%d completion=%d total=%d",
				responseObj.Usage.PromptTokens,
				responseObj.Usage.CompletionTokens,
				responseObj.Usage.TotalTokens,
			),
		)
	}

	log.Info.Printf(logString.String())

	response.Body = io.NopCloser(&body)
	return nil
}

func (s *Proxy) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	user := r.Header.Get("openai-api-proxy-user")
	requestID := r.Header.Get("opeanai-proxy-request-id")

	log.Error.Printf("error - %s - %s - %s - %s", requestID, r.URL.Path, user, err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// injecting user in the JSON body and updating the Content-Length header
func injectUser(r io.Reader, user string) (*bytes.Buffer, error) {
	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r).Decode(&bodyJSON)
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to decode request body"))
	}

	if bodyJSON["user"] == nil {
		bodyJSON["user"] = user
	}

	body, err := json.Marshal(bodyJSON)
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to encode request body"))
	}

	return bytes.NewBuffer(body), nil
}
