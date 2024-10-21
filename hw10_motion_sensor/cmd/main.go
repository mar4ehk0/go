package main

import (
	"fmt"
	"time"

	"github.com/mar4ehk0/hw10_motion_sensor/internal/calculator"
	"golang.org/x/exp/rand"
)

func main() {
	sensorCh := make(chan int)
	stopCh := make(chan struct{})
	dataCh := make(chan int)
	resultCh := make(chan float64)

	go readSensor(sensorCh, stopCh)

	go calculator.ProcessData(dataCh, resultCh)

LOOP:
	for {
		select {
		case <-stopCh:
			fmt.Println("\nStop read from sensor")
			break LOOP
		case sensorData := <-sensorCh:
			fmt.Print(".")
			dataCh <- sensorData
		case result := <-resultCh:
			fmt.Printf("\nAverage %f\n", result)
		}
	}
}

func readSensor(sensorCh chan<- int, stopCh chan<- struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout := time.After(15 * time.Second)

	for {
		select {
		case <-timeout:
			stopCh <- struct{}{}

			return
		case <-ticker.C:
			data := rand.Intn(100)
			sensorCh <- data
		}
	}
}
