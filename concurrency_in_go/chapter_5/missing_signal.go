package main

import (
	"fmt"
	"sync"
)

// child go routine
func doWork(cond *sync.Cond, done *bool) {

	fmt.Println("Work started !!")
	fmt.Println("Work Finished !!")

	// lock the signal before sending it
	cond.L.Lock()
	*done = true
	cond.Signal() // signals that the work is finished
	cond.L.Unlock()
}

func MainMissingSignal() {
	cond := sync.NewCond(&sync.Mutex{})

	cond.L.Lock()
	for i := 0; i < 50000; i++ {

		done := false

		go doWork(cond, &done)
		fmt.Println("Waiting for Child go routine...", i)
		// runtime.Gosched() // giving more chance to child go routine to execute
		// waits for the signal to continue

		if !done {
			cond.Wait()
		}

		fmt.Println("child goroutine finished..", i)
	}

	cond.L.Unlock()

}
