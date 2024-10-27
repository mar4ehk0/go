package main

import (
	"log"
	"time"

	"github.com/mar4ehk0/go/hw10_motion_sensor/internal/calculator"
	"golang.org/x/exp/rand"
)

func readSensor(sensorCh chan<- int, timeout <-chan time.Time) {
	defer close(sensorCh)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// start := time.Now()
	// done := make(chan bool)
	// defer close(done)
	// go func() {
	// 	<-timeout
	// 	log.Println("Прошло времени1:", time.Since(start))
	// 	done <- true
	// }()

	for {
		<-ticker.C
		data := rand.Intn(100)
		select {
		case sensorCh <- data:
		case <-timeout:
			return
		}
	}
}

func main() {
	sensorCh := make(chan int)
	averageCh := make(chan float64)

	timeout := time.After(60 * time.Second)

	go readSensor(sensorCh, timeout)

	go calculator.Average(sensorCh, averageCh)

	for x := range averageCh {
		log.Println("average: ", x)
	}
}
