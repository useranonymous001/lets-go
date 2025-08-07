package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounter() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	// 50 roroutines spawned
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// each goroutine loops for 1000 operation
			// performs 50000 operations
			for range 1000 {
				ops.Add(1)
				// ops++  // this in invalid operation
			}
		}()
	}
	wg.Wait()

	// no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a value
	// even while other goroutines are (atomically) updating it.
	fmt.Println("ops: ", ops.Load())
}
