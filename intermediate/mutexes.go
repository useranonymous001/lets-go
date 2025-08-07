package main

import (
	"fmt"
	"sync"
)

type Container struct {
	counters map[string]int
	mu       sync.Mutex
}

// Note: the mutex should not be copied, so when we pass around contaier
// we need to pass the pointer

func (c *Container) inc(name string) {
	// modifying the structure
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func MainMutexes() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {

		defer wg.Done()
		for range n {
			c.inc(name)
		}
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println("Counters: ", c.counters)
}

func Test() {

}
