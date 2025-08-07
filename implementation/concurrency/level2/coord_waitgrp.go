/*
4. Wait for N Workers to Finish
Spin off 10 goroutines that do some "work" (e.g., sleep + print).
The main goroutine should wait until all of them
are done before exiting.

✅ Use sync.WaitGroup
✅ Learn how Add, Done, and Wait work
*/

/*
A Waitgroup waits for all the goroutines to finish

Waitgroup.Add(int) => adds the number of goroutine to wait for
	- Needs to be added before any statement creating the goroutine

Waitgroup.Done() => signals after the goroutine is completed
	- actually this decrement the counter of goroutine that was incremented by using .Add() method

Waitgroup.Wait() => waits until all the goroutine throws Done()

*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// lets make a program that gets the contents of 20 webpages
// using 20 different goroutines

func MainWaitGroup() {

	wg := sync.WaitGroup{}
	for i := 30; i <= 50; i++ {

		url := fmt.Sprintf("https://www.ietf.org/rfc/rfc10%d.txt", i)
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("eerr")
			}
			defer resp.Body.Close()

			bytes, _ := io.ReadAll(resp.Body)

			fmt.Println(len(bytes))

		}(url)

	}
	fmt.Println("waiting....")
	wg.Wait()
	fmt.Println("Completed...")
}
