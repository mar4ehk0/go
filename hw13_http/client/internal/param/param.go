package param

import (
	"github.com/spf13/pflag"
)

type InputParam struct {
	Url    string
	Path   string
	Method string
	Body   string
}

func ReadFlags() *InputParam {
	url := pflag.StringP("url", "u", "localhost", "URL with port to connect for server")
	path := pflag.StringP("path", "p", "posts", "end-point to connect")
	method := pflag.StringP("method", "m", "GET", "HTTP method")
	body := pflag.StringP("body", "b", "", "HTTP body ALLOWED for post method")

	pflag.Parse()

	return &InputParam{Url: *url, Path: *path, Method: *method, Body: *body}
}
