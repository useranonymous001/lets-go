/*
6. Reader-Writer Lock
Simulate multiple readers and writers on a shared log buffer.

Multiple readers can read at once.
Only one writer can write at a time.
No reader can read while writing is in progress.

✅ Use sync.RWMutex
✅ Show how reads and writes interleave safely

*/

/*
Reader-Writer Lock
- Allows single writer but multiple reader

*/

package main

import (
	"fmt"
	"sync"
)

func MainRealisticCoord() {
	mutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	log := make([]string, 0, 10) // []

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Reader(&log, &mutex)
		}()
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Writer(&log, fmt.Sprintf("Item %d", i), &mutex)
		}()
	}

	wg.Wait()

}

func Reader(log *[]string, mutex *sync.RWMutex) {
	fmt.Println("reading....")
	mutex.RLock()
	for _, l := range *log {
		fmt.Print(l)
	}
	mutex.RUnlock()

}

func Writer(log *[]string, item string, mutex *sync.RWMutex) {

	fmt.Println("...writing")
	mutex.Lock()
	*log = append(*log, item)
	mutex.Unlock()
}
