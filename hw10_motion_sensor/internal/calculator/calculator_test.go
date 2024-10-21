package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanProcessDataWhenMore10Items(t *testing.T) {
	dataCh := make(chan int)
	resultCh := make(chan float64)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expected := 5.5

	go func() {
		for _, v := range numbers {
			dataCh <- v
		}
	}()

	go ProcessData(dataCh, resultCh)

	actual := <-resultCh

	assert.Equal(t, expected, actual)
}

func TestCanProcessDataWhenLess10Items(t *testing.T) {
	dataCh := make(chan int)
	resultCh := make(chan float64)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	expected := 4

	go func() {
		for _, v := range numbers {
			dataCh <- v
		}
	}()

	go ProcessData(dataCh, resultCh)

	var actual float64
	go func() {
		actual = <-resultCh
	}()

	assert.NotEqual(t, expected, actual)
}
