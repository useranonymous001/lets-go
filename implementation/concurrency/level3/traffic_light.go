/*
7. Traffic Light Controller
You have 3 lanes: red, yellow, green. Only one can be active at a time.

Change light every 2 seconds in a loop.
Multiple goroutines may be waiting for green to proceed.
Ensure only green allows passing.

✅ Use sync.Cond for coordination
✅ Use shared state to represent the current light
*/

/*

Light: green
 pass => goroutine


Light red
	red: Waiting for green
Light Yellow
	yellow: waiting for green
Light Green
	green: proceed
*/

/*

Here's multiple goroutines are said to be waited. So, if i used Signal() , it would only wake only one
waiting goroutine.
Since, multiple go routines would be waiting for the signal to be green. i need to use broadcast.


*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var maxAllowedCars = 3

func MainTrafficLightController() {

	// mutex := sync.Mutex{}
	cond := sync.NewCond(&sync.Mutex{})
	lights := []string{"red", "green", "yellow"}

	currLight := "red"
	sharedLight := &currLight

	passed := 0

	// spawn multiple goroutine
	for i := 0; i < 10; i++ {
		go Traffic(i, sharedLight, &passed, cond)
	}

	// THE BELOW GIVEN CODE IS JUST A PIECE OF SHIT
	// IT IS LOCKING THE VARIABLE FOR 6 STRAIGHT SECONDS

	// cond.L.Lock()
	// for _, light := range lights {
	// 	currLight = light
	// 	cond.Signal()
	// 	time.Sleep(2 * time.Second)
	// }
	// cond.L.Unlock()

	// INSTEAD FOLLOW THE FOLLOWING THINGS
	for {

		for _, light := range lights {

			cond.L.Lock()
			*sharedLight = light

			fmt.Println("Light changed to ", light)

			if light == "green" {
				passed = 0
			}

			// after changing the light, broadcast to all the vehicles(goroutine) those who are waiting
			// for light to change

			// cond.Signal() // this only wakes the single goroutine
			// i have to wake all the goroutine
			cond.Broadcast()
			cond.L.Unlock()
			// wait for 2 second to change the light
			time.Sleep(time.Second * 2)
		}

	}

}

func Traffic(id int, light *string, passed *int, cond *sync.Cond) {

	cond.L.Lock()
	for *light != "green" || *passed >= maxAllowedCars {
		cond.Wait()
	}
	fmt.Printf("🚗 Car %d is going on green light!\n", id)
	*passed++
	cond.L.Unlock()
	time.Sleep(1000 * time.Millisecond)

}
