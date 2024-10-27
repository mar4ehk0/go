package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAverage(t *testing.T) {
	sensorCh := make(chan int)
	averageCh := make(chan float64)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expected := 5.5

	go func() {
		for _, v := range numbers {
			sensorCh <- v
		}

		close(sensorCh)
	}()

	go Average(sensorCh, averageCh)

	actual := <-averageCh

	assert.Equal(t, expected, actual)
}

func TestCanAverageWhenLess10Items(t *testing.T) {
	sensorCh := make(chan int)
	averageCh := make(chan float64)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	expected := 0.0

	go func() {
		for _, v := range numbers {
			sensorCh <- v
		}

		close(sensorCh)
	}()

	go Average(sensorCh, averageCh)

	actual := <-averageCh

	assert.Equal(t, expected, actual)
}
