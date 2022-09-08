package main

import "sync"

type Counter struct {
	value int
	mx    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mx.Lock()
	c.value++
	c.mx.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
