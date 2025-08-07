/*
- Instead of using explicit locking the mutex to synchronize the access to the shared memory.

- Another approach is to use the built-in features for synchronization of shared access using
go routines and channels that also aligns with the Go's idea: share memory by communicating.

- Here, the state is owned by one goroutine and other goroutines have to send messages to the owing
goroutine inorder to get the corresponding replies.

- This makes sure that data is never corrupt with concurrent access.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func MainShareByComm() {

	// counter variables for counting the number of operations performed
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)   // reads channel that accepts type readOp
	writes := make(chan writeOp) // writes channel that accepts type writeOp

	// main goroutine that owes the main state
	go func() {
		state := map[int]int{}
		// this func either receives read request or write request
		// so handle it efficiently
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // send the current value of the struct
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	// simulating multiple read operations
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// simulating write operations
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}

				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	finalReadOps := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", finalReadOps)
	finalWriteOps := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", finalWriteOps)

}
