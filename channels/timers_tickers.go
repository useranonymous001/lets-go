package main

import (
	"fmt"
	"time"
)

// timers, execute go code at some point or at certain interval of time.

func MainTimer() {

	// NewTimer creates a new Timer that will send the current time
	// on its channel after at least duration d.
	timer1 := time.NewTimer(time.Second * 1)
	// defer timer1.Stop()
	// ok := timer1.Reset(time.Second * 2)
	<-timer1.C // blocks the timers channel C until it sends the value indicating timer is fired
	fmt.Println("Timer 1 fired")

	// you can also stop the timer, before it even get fired
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C // blocking the timer2 channel until it gets the value
		fmt.Println("Timer 2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(time.Second * 2)
}

func MainTicker() {

	// will send the current time on the channel after each tick.
	ticker := time.NewTicker(300 * time.Millisecond) // after each 500 milisecond send me a tick through channel
	done := make(chan bool)

	go func() {

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println(t)
			}
		}

	}()

	time.Sleep(time.Millisecond * 2500)
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
