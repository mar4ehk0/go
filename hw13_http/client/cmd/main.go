package main

import (
	"fmt"
	"os"

	myHttp "github.com/ma4ehk0/go/hw13_http/client/internal/http"
	"github.com/ma4ehk0/go/hw13_http/client/internal/param"
)

type InputParam struct {
	Url    string
	Path   string
	Method string
	Body   string
}

func main() {
	param := param.ReadFlags()

	httpRequest, err := myHttp.NewHttpRequest(*param)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rsp, err := myHttp.SendRequest(*httpRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.Write(rsp)
}
