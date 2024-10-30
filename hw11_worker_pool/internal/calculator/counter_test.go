package calculator

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanInc(t *testing.T) {
	counter := Counter{Mutex: sync.Mutex{}}
	expected := 2

	counter.Inc()
	counter.Inc()

	assert.Equal(t, expected, counter.Value)
}

func TestCanIncConcurrent(t *testing.T) {
	counter := Counter{Mutex: sync.Mutex{}}
	expected := 3

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	assert.Equal(t, expected, counter.Value)
}
