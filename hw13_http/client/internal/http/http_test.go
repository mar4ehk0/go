package http

import (
	"testing"

	"github.com/ma4ehk0/go/hw13_http/client/internal/param"
	"github.com/stretchr/testify/assert"
)

func TestCanNewHttpRequest(t *testing.T) {
	tests := []struct {
		name     string
		param    param.InputParam
		expected HttpRequest
	}{
		{
			"NewHttpRequest GET",
			param.InputParam{Url: "example.com", Path: "urn", Method: "GET", Body: ""},
			HttpRequest{Url: "http://example.com/urn", Method: "GET", Body: ""},
		},
		{
			"NewHttpRequest POST",
			param.InputParam{Url: "example.com", Path: "urn", Method: "POST", Body: "lorem ipsum"},
			HttpRequest{Url: "http://example.com/urn", Method: "POST", Body: "lorem ipsum"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := NewHttpRequest(tc.param)

			assert.Equal(t, tc.expected.Url, actual.Url)
			assert.Equal(t, tc.expected.Method, actual.Method)
			assert.Equal(t, tc.expected.Body, actual.Body)
		})
	}
}

func TestFailNewHttpRequest(t *testing.T) {
	tests := []struct {
		name     string
		param    param.InputParam
		expected error
	}{
		{
			"NewHttpRequest empty URL",
			param.InputParam{Url: "", Path: "urn", Method: "GET", Body: ""},
			ErrUrlEmpty,
		},
		{
			"NewHttpRequest empty Path",
			param.InputParam{Url: "example.com", Path: "", Method: "POST", Body: "lorem ipsum"},
			ErrPathEmpty,
		},
		{
			"NewHttpRequest empty Method",
			param.InputParam{Url: "example.com", Path: "urn", Method: "", Body: ""},
			ErrHttpMethodEmpty,
		},
		{
			"NewHttpRequest not allowed method PUT",
			param.InputParam{Url: "example.com", Path: "urn", Method: "PUT", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest not allowed method PATCH",
			param.InputParam{Url: "example.com", Path: "urn", Method: "PATCH", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest not allowed method DELETE",
			param.InputParam{Url: "example.com", Path: "urn", Method: "DELETE", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest GET and not empty Body",
			param.InputParam{Url: "example.com", Path: "urn", Method: "GET", Body: "lorem ipsum"},
			ErrWrongHttpBody,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewHttpRequest(tc.param)

			assert.ErrorIs(t, tc.expected, err)
		})
	}
}
