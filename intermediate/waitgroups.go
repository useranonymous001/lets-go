package main

import (
	"fmt"
	"sync"
	"time"
)

func dealer(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func MainWorkGroup() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			dealer(i)
		}()
	}

	wg.Wait()

	// this part is blocked by the wait function
	fmt.Println("waiting..")

}
