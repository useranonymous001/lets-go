/*
5. One-Time Initialization (Once)
Use sync.Once to initialize a configuration (e.g., load from file).
Simulate 10 goroutines trying to access the config — only one should trigger the load function.

✅ Use sync.Once
✅ Confirm config load message prints once
*/

/*

sync.Once

type Once struct {
}

- Once is an object that will perform exactly one action.
- A Once must not be copied after first use.
- In the terminology of the Go memory model,
- the return from f “synchronizes before” the return from any call of once.Do(f).

- Even you call multiple once.Do(func), the first call will only get invoked even if you have different instance in each call

- For that; a new instance of "Once" is required for each function

- Do is intended for initialization that must be run exactly once. Since f is niladic, it may be necessary to use a function literal to capture the arguments to a function to be invoked by Do:


*/

package main

import (
	"fmt"
	"sync"
)

func MainOneTimeInit() {

	once := sync.Once{}

	work := func() {
		fmt.Println("Did Once")
	}

	channel := make(chan bool) // unbuffered channel
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(work)
			// sender
			channel <- true
		}()
	}

	// either you use wait for 10 goroutine or 1, the first fn call only gets executed
	for i := 0; i < 10; i++ {
		<-channel // receivers // discard the data; ignore
	}

	// why use channel
	// channels blocks the execution, until the sender sends the signal/message
	// Receivers always block until there is data to receive. If the channel is unbuffered, the sender blocks until the receiver has received the value.

	// If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.

}
