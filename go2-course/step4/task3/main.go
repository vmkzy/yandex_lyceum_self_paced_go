package main

import "sync"

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = append(c.queue, element)
}
func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.queue) == 0 {
		return nil
	}
	res := c.queue[0]
	c.queue = c.queue[1:]
	return res
}
