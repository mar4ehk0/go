package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ma4ehk0/go/hw13_http/client/internal/param"
)

type Request struct {
	URL    string
	Method string
	Body   string
}

var (
	ErrURLEmpty        = errors.New("url server empty")
	ErrPathEmpty       = errors.New("path server empty")
	ErrHTTPMethodEmpty = errors.New("http method empty")
	ErrWrongHTTPBody   = errors.New("body does not allowed for GET method")
	ErrNotImplMethod   = errors.New("not implemented yet method")
)

const (
	HTTPGet  string = "GET"
	HTTPPost string = "POST"
)

func NewHTTPRequest(param param.InputParam) (*Request, error) {
	http := &Request{}

	if len(param.URL) == 0 {
		return nil, ErrURLEmpty
	}
	if len(param.Path) == 0 {
		return nil, ErrPathEmpty
	}
	if len(param.Method) == 0 {
		return nil, ErrHTTPMethodEmpty
	}
	if !(param.Method == HTTPGet || param.Method == HTTPPost) {
		return nil, ErrNotImplMethod
	}

	if param.Method == HTTPGet && len(param.Body) != 0 {
		return nil, ErrWrongHTTPBody
	}

	url := fmt.Sprintf("http://%s/%s", param.URL, param.Path)

	http.URL = url
	http.Method = param.Method
	http.Body = param.Body

	return http, nil
}

func SendRequest(request Request) ([]byte, error) {
	var processedResp []byte
	var err error

	switch request.Method {
	case HTTPGet:
		processedResp, err = get(request.URL)
	case HTTPPost:
		processedResp, err = post(request.URL, request.Body)
	default:
		return nil, ErrNotImplMethod
	}

	return processedResp, err
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		return nil, fmt.Errorf("can't do get: %w", err)
	}

	processedResp, err := processResponse(resp)

	return processedResp, err
}

func post(url string, data string) ([]byte, error) {
	requestBody := []byte(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody)) //nolint
	if err != nil {
		return nil, fmt.Errorf("can't do post: %w", err)
	}

	processedResp, err := processResponse(resp)

	return processedResp, err
}

func processResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't process request: %w", err)
	}

	return body, nil
}
