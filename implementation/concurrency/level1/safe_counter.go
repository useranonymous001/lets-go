/*
1. Safe Counter with Mutex
Create a counter that can be safely incremented by 100 goroutines. Each goroutine should increment the counter 1000 times.

✅ Use sync.Mutex
✅ Print the final count (should be 100000)
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Question 1
	// MainSafeCounter()

	// Question 2
	// MainBankSimulation()

	// Question 3
	MainProdConsumer()
}

func MainSafeCounter() {

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	count := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter(&count, &mutex)
		}()

	}
	// time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("Counter: ", count)

}

func counter(count *int, mutex *sync.Mutex) {

	mutex.Lock()
	for i := 0; i < 1000; i++ {
		*count++
	}
	mutex.Unlock()
}
