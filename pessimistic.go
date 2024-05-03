package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func TestPessimistic() {
	var wg sync.WaitGroup
	var c counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				c.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value())
}
