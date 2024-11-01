package calculator

import (
	"strconv"
	"sync"
)

type Counter struct {
	Mutex sync.RWMutex
	Value int
}

func (c *Counter) Inc() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Value++
}

func (c *Counter) String() string {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	return strconv.Itoa(c.Value)
}
