package main

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

type Count interface {
	Increment()
	GetValue() int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	return c.value
}
