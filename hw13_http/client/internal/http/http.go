package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrNotImplMethod = errors.New("not implemented yet method")
)

func SendRequest(request HttpRequest) ([]byte, error) {
	var processedResp []byte
	var err error

	switch request.Method {
	case HttpGet:
		processedResp, err = get(request.Url)
	case HttpPost:
		processedResp, err = post(request.Url, request.Data)
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
