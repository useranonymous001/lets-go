package main

import (
	"fmt"
	"time"
)

func MainTimeOut() {

	c1 := make(chan string, 1)

	go func() {
		// perform some external work
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)

	// After waits for the duration to elapse and
	//  then sends the current time on the returned channel.
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		// performs some tasks (external task)
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	// this will be executed and doesn't block the execution
	default:
		fmt.Println("hello")
	}

}
