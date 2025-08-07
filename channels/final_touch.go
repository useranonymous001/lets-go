package main

import "fmt"

func MainFinalTouch() {

	// chan2 := make(chan string)
	// chan1 := make(chan string)

	// go func() {
	// 	chan1 <- "msg 1"
	// 	time.Sleep(time.Second * 1)
	// }()

	// go func() {
	// 	chan2 <- "msg 2"
	// 	time.Sleep(time.Second * 2)
	// }()

	// // main goroutine to consume data from channel
	// for range 2 {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println("msg1 received: ", msg1)

	// 	case msg2 := <-chan2:
	// 		fmt.Println("msg1 received: ", msg2)

	// 	}
	// }

	queue := make(chan string, 2)
	queue <- "go.sh"
	queue <- "script.sh"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
