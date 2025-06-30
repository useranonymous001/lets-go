package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func MainWeiredGoRoutines() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go callHello()
		go increment()
	}
	wg.Wait()
}

func callHello() {
	fmt.Printf("Hello #%d\n", counter)
	wg.Done() // decrements the number of go routine spawned
}

func increment() {
	counter++
	wg.Done()
}

/*
Output:

Hello #1
Hello #1
Hello #2
Hello #3
Hello #4
Hello #5
Hello #8
Hello #9
Hello #9
Hello #7


It seems that the result/output is not being syncronized,
this is because, when the multiple go routines are spawned,
all of them competes against each other about who executes first

this is why the results are not synchronized to each other
*/
