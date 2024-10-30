package calculator

import (
	"strconv"
	"sync"
)

type Counter struct {
	Mutex sync.Mutex
	Value int
}

func (c *Counter) Inc() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Value++
}

func (c *Counter) String() string {
	return strconv.Itoa(c.Value)
}
