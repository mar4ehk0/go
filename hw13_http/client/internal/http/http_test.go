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
		expected Request
	}{
		{
			"NewHttpRequest GET",
			param.InputParam{URL: "example.com", Path: "urn", Method: "GET", Body: ""},
			Request{URL: "http://example.com/urn", Method: "GET", Body: ""},
		},
		{
			"NewHttpRequest POST",
			param.InputParam{URL: "example.com", Path: "urn", Method: "POST", Body: "lorem ipsum"},
			Request{URL: "http://example.com/urn", Method: "POST", Body: "lorem ipsum"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := NewHTTPRequest(tc.param)

			assert.Equal(t, tc.expected.URL, actual.URL)
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
			param.InputParam{URL: "", Path: "urn", Method: "GET", Body: ""},
			ErrURLEmpty,
		},
		{
			"NewHttpRequest empty Path",
			param.InputParam{URL: "example.com", Path: "", Method: "POST", Body: "lorem ipsum"},
			ErrPathEmpty,
		},
		{
			"NewHttpRequest empty Method",
			param.InputParam{URL: "example.com", Path: "urn", Method: "", Body: ""},
			ErrHTTPMethodEmpty,
		},
		{
			"NewHttpRequest not allowed method PUT",
			param.InputParam{URL: "example.com", Path: "urn", Method: "PUT", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest not allowed method PATCH",
			param.InputParam{URL: "example.com", Path: "urn", Method: "PATCH", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest not allowed method DELETE",
			param.InputParam{URL: "example.com", Path: "urn", Method: "DELETE", Body: ""},
			ErrNotImplMethod,
		},
		{
			"NewHttpRequest GET and not empty Body",
			param.InputParam{URL: "example.com", Path: "urn", Method: "GET", Body: "lorem ipsum"},
			ErrWrongHTTPBody,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewHTTPRequest(tc.param)

			assert.ErrorIs(t, tc.expected, err)
		})
	}
}
