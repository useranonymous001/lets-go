package main

import (
	"fmt"
	"sync"
)

// Go Routines

/*

Go Runtime is a part of the executable Binaries that is created while compiling GO Programs

Go Runtime Contains:
	- Garbage Collector
	- Scheduler to maintain light weight threads called GO ROUTINES

Go routine executes independently
	i.e., you can invoke any number of go routines and all of them could run concurrently

How multiple Go Routines Communicate to each other??
-> Via typed conduits called CHANNELS (can be used to send or receive data between go routines)

*/

/*
Go Routine is like a process that is running in the background
*/

/*
	sync.WorkGroup package


	var wg = sync.WorkGroup

	wg.Add(int)  => add the number of go routines to wait for
	wg.Wait()   => block the further operation until the go routines are completed execution
	wg.Done()   => indicates that the task has been completed of that go routine
	             and now decrement the number of go routine to wait for

*/

var msg string
var c = make(chan int)
var workGroup sync.WaitGroup

func setMessage() {
	msg = "This is Go Routine"
	// c <- 0 // sending the value to the receiver
	workGroup.Done()
}

func sayHello() {
	fmt.Println("Hello Gophers")
	workGroup.Done()
}

// main() is special go routine invoked during the startup of any program.
func main() {
	// workGroup.Add(2)
	// go setMessage() // called as a goroutine so, goes to the background and main goroutine exec continues
	// go sayHello()

	// time.Sleep(1 * time.Millisecond) // Sleep pauses the current goroutine(main go-routine for now) for at least the duration d
	// workGroup.Wait()

	// fmt.Println(msg)

	// other funcs
	// channelInGo()
	// channelInWork()

	// MainGoRoutine()
	// MainWeiredGoRoutines()
	MainMutexes()
}

// CHANNELS

/*
Like, we discussed before, channels are used to communicate between multiple go routines
They can be used to send and receive any kind of values using CHANNEL OPERATOR <-

channel of int values
ch := make(chan int)
*/

func channelInGo() {

	ch := make(chan int)

	// send value to the channel
	ch <- 4

	// receive value from the channel and assign to the variables
	v := <-ch

	fmt.Println(v)

	// this func is only for the demonstration of how data are exchanged through the channels, this throws fatal error of deadlock
}

func channelInWork() {

	go setMessage()
	v := <-c // it blocks the execution until it receives the data from its go routine
	fmt.Println(msg, v)

}
