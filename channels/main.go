package main

// var ch = make(chan string)

func main() {
	// spawner()
	// for i := 0; i < 5; i++ {
	// 	// receiving msg from the channel
	// 	fmt.Println(<-ch)
	// }
	// MainChannl()
	// MainSelect()

	// MainForSelect()
	// MainSelectDone()
	// MainFinalTouch()
	// MainTimeOut()
	// MainPingPong()
	// MainTimer()
	MainTicker()
}

// func spawner() {
// 	for i := 0; i < 5; i++ {
// 		go func(n int) {
// 			// sending msg to another goroutine using channels.
// 			ch <- fmt.Sprintf("From Go Routine: %d", n)
// 		}(i)
// 	}
// }
