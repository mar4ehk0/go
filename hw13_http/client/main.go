package main

import (
	"fmt"
	"os"

	myHttp "github.com/ma4ehk0/go/hw13_http/client/internal/http"
)

func main() {
	httpRequest, err := myHttp.NewHttpRequest()
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
