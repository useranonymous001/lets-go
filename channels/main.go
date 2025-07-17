package main

import "fmt"

var ch = make(chan string)

func main() {
	spawner()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func spawner() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			ch <- fmt.Sprintf("From Go Routine: %d", n)
		}(i)
	}
}
