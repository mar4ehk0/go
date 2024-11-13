package main

import (
	"fmt"
	"os"

	"github.com/mar4ehk0/go/hw12_log_util/internal/analyzer"
	"github.com/mar4ehk0/go/hw12_log_util/internal/file"
	"github.com/mar4ehk0/go/hw12_log_util/internal/param"
)

func main() {
	fileToPath, method, output := param.GetInputParam()

	stat := analyzer.NewStat(method)

	out := make(chan string)
	go func() {
		err := file.ReadFile(fileToPath, out)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	for v := range out {
		stat.Analyze(v)
	}

	if len(output) == 0 {
		fmt.Println(stat)
	} else {
		err := file.CreateFile(output, *stat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
