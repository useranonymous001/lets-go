package main

import (
	"fmt"
)

func MainSelect() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "hey i am channel 1"
	}()

	go func() {
		channel2 <- "hey i am channel 2"
	}()

	// fmt.Println("asdf: ", <-channel1)

	// blocks the code until anyone of the channel recieves the message.
	// if multiple channels are receving the msg, any one who get the first
	// would be executed.
	select {
	case msgFromChannel1 := <-channel1:
		fmt.Println(msgFromChannel1)
	case msgFromChannel2 := <-channel2:
		fmt.Println(msgFromChannel2)
	}
}
