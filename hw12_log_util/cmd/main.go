package main

import (
	"fmt"
	"os"

	"github.com/mar4ehk0/go/hw12_log_util/internal/analyzer"
	"github.com/mar4ehk0/go/hw12_log_util/internal/file"
	"github.com/mar4ehk0/go/hw12_log_util/internal/input_param"
)

func main() {

	fileToPath, method, output := input_param.GetInputParam()

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

// package main

// import (
// 	"bufio"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"
// )

// const logfile = "../access.log"

// func scanFile() error {
// 	total := 0 // count lines
// 	begin := time.Now()

// 	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
// 	if err != nil {
// 		log.Fatalf("open file error: %v", err)
// 		return err
// 	}
// 	defer f.Close()

// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() {
// 		a := sc.Text()

// 		a = a
// 		total++
// 	}
// 	if err := sc.Err(); err != nil {
// 		log.Fatalf("scan file error: %v", err)
// 		return err
// 	}

// 	log.Printf("scan file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)

// 	return nil
// }

// func readFileLines() error {
// 	total := 0
// 	begin := time.Now()

// 	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
// 	if err != nil {
// 		log.Fatalf("open file error: %v", err)
// 		return err
// 	}
// 	defer f.Close()
// 	rd := bufio.NewReader(f)
// 	for {
// 		if _, err := rd.ReadString('\n'); err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			log.Fatalf("read file line error: %v", err)
// 			return err
// 		}
// 		total++
// 	}

// 	log.Printf("reader read string in file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
// 	return nil
// }

// func readFileOnce() error {
// 	total := 0
// 	begin := time.Now()

// 	data, err := ioutil.ReadFile(logfile)
// 	if err != nil {
// 		return err
// 	}
// 	ss := strings.Split(string(data), "\n")
// 	for _, s := range ss {
// 		_ = s
// 		total++
// 	}

// 	log.Printf("read file once and split strings, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
// 	return nil
// }

// func main() {
// 	scanFile()
// 	readFileLines()
// 	readFileOnce()
// }
