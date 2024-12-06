package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ma4ehk0/go/hw13_http/client/internal/param"
)

type HttpRequest struct {
	Url    string
	Method string
	Body   string
}

var (
	ErrUrlEmpty          = errors.New("url server empty")
	ErrPathEmpty            = errors.New("path server empty")
	ErrHttpMethodEmpty      = errors.New("http method empty")
	ErrWrongHttpBody        = errors.New("body does not allowed for GET method")
	ErrNotImplMethod        = errors.New("not implemented yet method")
)

const HttpGet string = "GET"
const HttpPost string = "POST"

func NewHttpRequest(param param.InputParam) (*HttpRequest, error) {

	http := &HttpRequest{}

	if len(param.Url) == 0 {
		return nil, ErrUrlEmpty
	}
	if len(param.Path) == 0 {
		return nil, ErrPathEmpty
	}
	if len(param.Method) == 0 {
		return nil, ErrHttpMethodEmpty
	}
	if !(param.Method == HttpGet || param.Method == HttpPost) {
		return nil, ErrNotImplMethod
	}

	if param.Method == HttpGet && len(param.Body) != 0 {
		return nil, ErrWrongHttpBody
	}

	url := fmt.Sprintf("http://%s/%s", param.Url, param.Path)

	http.Url = url
	http.Method = param.Method
	http.Body = param.Body

	return http, nil
}

func SendRequest(request HttpRequest) ([]byte, error) {
	var processedResp []byte
	var err error

	switch request.Method {
	case HttpGet:
		processedResp, err = get(request.Url)
	case HttpPost:
		processedResp, err = post(request.Url, request.Body)
	default:
		return nil, ErrNotImplMethod
	}

	return processedResp, err
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		err := fmt.Errorf("can't do get: %w", err)

		return nil, err
	}

	processedResp, err := processResponse(resp)

	return processedResp, err
}

func post(url string, data string) ([]byte, error) {
	requestBody := []byte(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		err := fmt.Errorf("can't do post: %w", err)

		return nil, err
	}

	processedResp, err := processResponse(resp)

	return processedResp, err
}

func processResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("can't process request: %w", err)

		return nil, err
	}

	return body, nil
}
