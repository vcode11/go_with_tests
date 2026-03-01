package sync_counter

import "sync"



type Counter struct {
	mu sync.Mutex
	val int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val += 1
}

func (c *Counter) Value() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

