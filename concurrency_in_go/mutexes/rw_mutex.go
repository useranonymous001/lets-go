package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func MainRwMutex() {

	// mutex := sync.Mutex{}
	mutex := sync.RWMutex{} // initializing the read write locker
	events := make([]string, 0, 10000)

	// pre populating the events
	// simulating the game was already on going with some events
	// recoreded previously
	for i := 0; i < 10000; i++ {
		events = append(events, "Match Event")
	}

	// calling the match recorder to record the match events
	// then sending it to the client (multiple clients)
	go MatchRecorder(&events, &mutex)

	start := time.Now()
	for i := 0; i < 5000; i++ {
		// simulating large clients reading the events
		go ClientHandler(&events, &mutex, start)
	}

	time.Sleep(100 * time.Second)
}

// The matchRecorder requires Lock() and Unlock(), since it
// needs to update the data

// func MatchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
func MatchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, "Match Event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Append match event")
	}
}

// the clienthandler () requires RLock() and RUnlock() cause
// it is just reading the data

// ######

// *** Why Read Locks are required? ***
// cause we don't want to modify the contents of the slice
// while we are traversing the data
// modifying the pointer and contents of the slice while another goroutine is travers-
// ing it might lead us to follow an invalid pointer reference.

// ######

// func ClientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
func ClientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {

	// copy the entire match events simulating the response to the client
	// acting as copying the events from the match recoreder and sending to client concurrently
	for i := 0; i < 100; i++ {
		// mutex.Lock()
		mutex.RLock()
		allEvents := CopyAllEvents(mEvents)
		// mutex.Unlock()
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Printf("Copied %d events in %v\n", len(allEvents), timeTaken)
	}
}

func CopyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, len(*matchEvents))
	// for _, e := range *matchEvents {
	// You preallocate a slice of len(...) and then append to it, which doubles the slice length.
	// allEvents = append(allEvents, e)
	copy(allEvents, *matchEvents)
	// }
	return allEvents

}
