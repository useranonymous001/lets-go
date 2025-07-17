package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func MainMutexes() {
	mutex := sync.Mutex{}
	frequency := make([]int, 26)

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}

	// let's monitor the frequency table
	// tryLock() does not block the execution but checks whether the
	// resources is locked or not
	// if not locked, does the operations (generally reading) by lockign
	// then unlocks after the operations performed.

	for i := 0; i < 100; i++ {
		// sleep for 100 ms and try again locking
		// with tryLock()
		time.Sleep(100 * time.Millisecond)

		if mutex.TryLock() {
			for cIndex, char := range letters {
				fmt.Printf("%c-%d ", char, frequency[cIndex])
			}
			mutex.Unlock()
		} else {
			fmt.Println("Mutex already being used")
		}
	}

}

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	// download the document
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error downloading document..")
	}
	defer resp.Body.Close()

	// process the document
	bytes, _ := io.ReadAll(resp.Body)
	// bytes = [01101011, 01001110, 11010101]

	mutex.Lock()
	for _, byte := range bytes {
		char := strings.ToLower(string(byte))
		cIndex := strings.Index(letters, char) // returns the position of the substring char in strings letters if available
		// update the frequency counter
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	mutex.Unlock()
	fmt.Println("Completed: ", url, time.Now().Format("15:04:05"))

}
