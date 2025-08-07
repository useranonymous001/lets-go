// Rate Limiter:
// Limits the number of access per time

package main

import (
	"fmt"
	"time"
)

func RateLimiter() {

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("")

	// we can also create a bursty rate limiter
	// that is, multiple request at a time but with limit
	burstyLimiter := make(chan time.Time, 3) // allowing 3 request at a time

	// fill the burstyLimiter already so that we can receive
	// This pre-fills the channel with 3 tokens, allowing 3 requests to pass immediately without waiting.
	// That’s the "burst" part — to allow an initial burst of traffic.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// This is the rate limiter. It gradually adds new tokens to the burstyLimiter,
	// simulating "one token every 200ms".
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req, time.Now())
	}

}
