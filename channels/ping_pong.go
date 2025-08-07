package main

import "fmt"

// sends message
// accepts one channel for sending msg
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// receives message
// accepts two channels:
// pings => receives
// pongs => sends
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func MainPingPong() {

	pings := make(chan string, 1) // buffered channel
	pongs := make(chan string, 1)
	ping(pings, "what the hell")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
