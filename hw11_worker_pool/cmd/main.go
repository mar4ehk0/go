package main

import (
	"fmt"
	"sync"

	"github.com/mar4ehk0/go/hw11_worker_pool/internal/calculator"
)

func main() {
	wg := sync.WaitGroup{}

	counter := calculator.Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			counter.Inc()
		}()
	}
	wg.Wait()

	fmt.Println(&counter)
}
