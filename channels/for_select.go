package main

import "fmt"

func MainForSelect() {

	channel := make(chan string, 3) // buffered channel with 3 size
	data := []string{"messi", "ronaldo", "neymar", "lewandoski"}

	// pass the data through the channels and loop through the data stored in the channels

	for _, d := range data {

		// select {
		// case channel <- d:
		// }
		channel <- d
	}

	for i := 0; i < 4; i++ {
		fmt.Println(<-channel)
	}

	// close the channel
	close(channel)

	// loop through the data stored in the channel
	for data := range channel {
		fmt.Println(data)
	}

}
