package http

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

type HttpRequest struct {
	Url    string
	Method string
	Data   string
}

var (
	ErrDomainEmpty          = errors.New("url server empty")
	ErrRouteEmpty           = errors.New("urn server empty")
	ErrHttpMethodEmpty      = errors.New("http method empty")
	ErrNotAllowedHttpMethod = errors.New("http method not allowed")
	ErrWrongHttpBody        = errors.New("body does not allowed for GET method")
)

const HttpGet string = "GET"
const HttpPost string = "POST"

func NewHttpRequest() (*HttpRequest, error) {
	domain, route, method, body := readFlags()

	http := &HttpRequest{}

	if len(domain) == 0 {
		return nil, ErrDomainEmpty
	}
	if len(route) == 0 {
		return nil, ErrRouteEmpty
	}
	if len(method) == 0 {
		return nil, ErrHttpMethodEmpty
	}
	if !(method == HttpGet || method == HttpPost) {
		return nil, ErrNotAllowedHttpMethod
	}

	if method == HttpGet && len(body) != 0 {
		return nil, ErrWrongHttpBody
	}

	url := fmt.Sprintf("http://%s/%s", domain, route)

	http.Url = url
	http.Method = method
	http.Data = body

	return http, nil
}

func readFlags() (string, string, string, string) {
	domain := pflag.StringP("domain", "d", "localhost", "Domain server to connect")
	route := pflag.StringP("route", "r", "home", "end-point to connect")
	method := pflag.StringP("method", "m", HttpGet, "HTTP method")
	body := pflag.StringP("body", "b", "", "HTTP body ALLOWED for post method")

	pflag.Parse()

	return *domain, *route, *method, *body
}
