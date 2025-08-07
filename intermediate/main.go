// workers pool
// how to implement worker pool using goroutine and channel

// pool of workers

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	// initially it is blocked, as jobs are not ready yet
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)

		results <- job * 2
	}

}

func main() {
	// WorkerMain()
	// MainWorkGroup()
	// RateLimiter()
	// AtomicCounter()
	// MainMutexes()
	// MainShareByComm()
	// MainPackage()
	// MainPanic()
	MainString()
}

func WorkerMain() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create a pool of worker
	// initially blocked as no jobs are provided
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// now send jobs to the channel
	// fuck i was just trying to send the number of jobs to be less than it should be
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// close the jobs channel, so it wouldn't send any data further.
	close(jobs)

	fmt.Println("after job result:")
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
