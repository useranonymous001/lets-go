package main

import (
	"fmt"
	"time"
)

func MainGoRoutine() {

	var msg = "hello"
	// it has now created a dependency issue between the
	// main go routine and the new go routine that is spinned off
	// the variable is being shared between the main go routine and the new go routine
	// this may create a race condition
	go func(msg string) { // de coupling
		fmt.Println(msg)
	}(msg)
	msg = "Gophers"
	time.Sleep(time.Millisecond * 100)
}
