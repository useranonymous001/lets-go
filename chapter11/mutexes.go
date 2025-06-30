package main

import (
	"fmt"
	"sync"
)

/*
*** Mutexes ***
Mutexes are basically locks, that an application is going to honor
*/

// shared resources
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	// lock the resources
	c.mu.Lock()
	// perform the critical memory task
	c.count++
	// unlock the resources
	c.mu.Unlock()
}

func (c *Counter) GetCount() int {

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count

}

func MainMutexes() {

	valueCounter := &Counter{}
	var wait_group = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wait_group.Add(1)
		go func() {
			valueCounter.Increment()
			wait_group.Done()
		}()
	}

	// don't let the main function exit before all the go routines executes
	wait_group.Wait()

	fmt.Println(valueCounter.GetCount())

}

/*
It is very cruicial to manage the accesing of same memory resources by synchronizing the
access time for each threads

this is called task synchronization
*/
