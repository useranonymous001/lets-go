// channels is used as pipe through which message are passed between multiple goroutines
// Main goroutine can easily access its child goroutines
// message are acced with FIFO approach
// receiving msg from the channel is synchronous. i.e., it block the execution of the program
// until the msg are received from the go routine.

package main

import (
	"fmt"
)

func MainChannl() {

	channel := make(chan string)

	go func() {
		// does some work
		channel <- "some data from anon goroutine"
		channel <- "msg2"

		// this doesn't gets chance to get printed
		// before this gets printed, the main goroutine gets executed and exits.
		// it may barely get its turn for execution.
		fmt.Println("i am go routine, sorry i am late")

	}()

	// main go routine
	// get the data from the go routine
	msg := <-channel // once the data is read from the channel
	// it cannot be consumed again another time

	// cannot receive same data.
	// blocks the execution of the code untile the msg is not received.
	msg2 := <-channel
	fmt.Println(msg, msg2)

	fmt.Println("what the fuckkk??")

	// if i try to wait for message, that's never going to be sent or is already consumed
	// it;s gonna be deadlock situation.
}
