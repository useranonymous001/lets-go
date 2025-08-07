package main

import "fmt"

// this is only the read-only channel
// this is why the done channel are usefull inside the for-select loop
// data can be read and that's what we need inside the for-select-loop

func MainSelectDone() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	// spawning a new go routine that works until the done message is sent

	go func() {
		for {
			// job - int || more - bool
			// more = true when job is actually received
			job, more := <-jobs
			if more {
				fmt.Println("job received, ", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()

	// pass jobs to the goroutine

	for i := 0; i < 3; i++ {
		jobs <- i + 1
		fmt.Println("job sent: ", i+1)
	}

	// close after finish sending jobs to job channel
	close(jobs)

	<-done

	fmt.Println("sent all jobs")

	_, ok := <-jobs
	fmt.Println("still job available??  ", ok)

}
