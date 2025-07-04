package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go doWork(i)
	}

	time.Sleep(2 * time.Second)
}

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05")) // Format returns a textual representation of the time value formatted according to the layout defined by the argument.
	time.Sleep(1 * time.Second)                                              // sleep for 1 second
	fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}
