/*
🧠 Level 4: Simulate System-like Behavior
8. Elevator Controller
Simulate an elevator system:

5 floors

Goroutines simulate passengers requesting floors

Elevator moves between floors, picking up and dropping off

Handle concurrent requests without deadlocks

✅ Use mutexes and condition variables to model movement and wait logic
*/

/*
Elelvator will be a single goroutine (loop)

passengers multiple goroutines (request Elevator goroutine)

*/

/*
This programs contains deadlock, since the elvator keep moving everytime and would be waiting for signals (request)
from the passengers, but we have very limited passengers. So, even after the elevator completes its tasks.
it keep moving and waiting for passenger, which is never there.
*/

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Request struct {
	from int
	to   int
}

/*
Explanation:

	** How elvator works?
	=> takes from and to
	=> drops at to(destination)

*/

// creating a shared slice of request
var requests []Request
var onBoard []Request
var wg sync.WaitGroup

func MainElevatorSystem() {
	// controller | request generator
	cond := sync.NewCond(&sync.Mutex{})
	maxFloors := 5
	floor := 0
	direction := 1 // 1 up 0 down

	wg.Add(2)

	go func() {

		defer wg.Done()

		for {
			// move the elevator
			floor += direction // if dir = 1, it moves upward

			if floor == maxFloors {
				direction = -1
			} else if floor == 0 {
				direction = 1
			}

			cond.L.Lock()
			if len(requests) == 0 && len(onBoard) == 0 {
				cond.Wait()
			}

			if len(requests) == 0 && len(onBoard) == 0 {
				os.Exit(1)
			}

			// pick the passengers
			newRequests := []Request{}
			for _, req := range requests {

				if req.from == floor {
					// pick up
					fmt.Printf("Picked up passenger from %d going to %d\n", req.from, req.to)
					onBoard = append(onBoard, req)
					time.Sleep(1 * time.Second)
				} else {
					// on-boarding the
					newRequests = append(newRequests, req)
				}
			}
			requests = newRequests

			newBoard := []Request{}
			for _, board := range onBoard {
				if board.to == floor {
					fmt.Printf("Dropped passenger from %d to %d\n", board.from, board.to)
					time.Sleep(time.Second * 1)
				} else {
					newBoard = append(newBoard, board)
				}
			}
			onBoard = newBoard
			cond.L.Unlock()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// passenger loops as much as they can
	// generating passenger as much we can
	go func() {
		defer wg.Done()
		cond.L.Lock()
		requests = append(requests, Request{
			from: 1,
			to:   3,
		}, Request{
			from: 3,
			to:   5,
		},
			Request{
				from: 4,
				to:   2,
			},
		)
		// now signal to the elevator that request is generated
		cond.Signal()
		cond.L.Unlock()
	}()

	wg.Wait()

}
