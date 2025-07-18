/*
3. Producer-Consumer (Bounded Buffer)
	Implement a producer-consumer system using a fixed-size slice (buffer).

		- Producer adds items to buffer.
		- Consumer removes items.
		- If buffer is full, producer waits.
		- If buffer is empty, consumer waits.

		✅ Use sync.Mutex + sync.Cond
		✅ Simulate multiple producers and consumers
*/

// slices.DeleteFunc() => search and delete the mathching item from the slice

package main

import (
	"fmt"
	"slices"
	"sync"
)

// recreating slice as a fixed size buffer.
func MainProdConsumer() {

	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	buffer := make([]int, 0, 10)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			producer(&buffer, i+10, cond)
		}()

	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			consumer(&buffer, i+10, cond)
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("go routine completed")

	cond.L.Lock()
	fmt.Printf("Buffer Len: %d\nBuffer: %v\n", len(buffer), buffer)
	cond.L.Unlock()
}

func producer(buffer *[]int, item int, cond *sync.Cond) {

	cond.L.Lock()

	for len(*buffer) == 10 {
		cond.Wait()
	}

	*buffer = append(*buffer, item)
	cond.Signal()
	fmt.Println("Producing...")

	cond.L.Unlock()

}

func consumer(buffer *[]int, item int, cond *sync.Cond) {

	// either delete the item from the index or item value itself
	// here's i am trying to delete the item with its value itself.
	cond.L.Lock()

	for len(*buffer) == 0 {
		cond.Wait()
	}

	*buffer = slices.DeleteFunc(*buffer, func(value int) bool {
		return item == value
	})
	cond.Signal()
	fmt.Println("Consuming...")
	cond.L.Unlock()

}
