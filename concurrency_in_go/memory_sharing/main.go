package main

import (
	"fmt"
	"sync"
	"time"
)

// main go routine reads the content from the memory
func main() {

	// protect the variable with mutex
	mutex := sync.Mutex{}
	count := 5                   // create a memory variable
	go countdown(&count, &mutex) // share the memory pointer reference to the go routine

	// check the value for count every half second
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}

	// MainUpdate()

}

// another go routine updates the content of the shared memory
// updating the shared memory every second
func countdown(seconds *int, mutex *sync.Mutex) {
	// seconds is not escaping to heap, and staying to the stack.
	// However the compiler is placing the count variable to the heap, as we are sharing this variable between the goroutines

	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		// mutex.Lock()

		if mutex.TryLock() {
			*seconds -= 1
			mutex.Unlock()
		}
	}
}

// if we would remove the go keyword from the countdown func
// the program would be then sequential.
// means that the loop inside the main stack wouldn't get executed

// There's some small additional cost of using heap memory instead of using stack memory.
// Cuz when we are done using the memory, the heap needs to be cleaned up by go's garbage collector.

// We can use a tool to analyze that a variable has escape to the heap by:
// - go tool compile -m go_file.go

// Inlining Function
// - This is a optimization technique of the compiler, where under certain conditions,
// the compiler replaces function call with the actual contents of the functioon itself.
