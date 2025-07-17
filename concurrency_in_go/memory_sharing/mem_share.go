// function producing letter frequency

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

func MainUpdate() {

	mutex := sync.Mutex{}

	frequency := make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	time.Sleep(10 * time.Second)

	mutex.Lock()
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	mutex.Unlock()
}

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {

	// locking the entire execution making it sequential
	// mutex.Lock()
	// instead just lock where its necessary
	// we are trying to lock all the process from downloading to modifying.
	// even if we try to just lock the modifying code that is updaing the frequency
	// it is also cause performance cost cuz calling lock and unlock too often
	// is also not good

	// so let's make the download process concurrent as it is.
	// then make the processing part sequential
	// i.e., files will be downloaded concurrently but the processing part will be sequential

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error while getting url")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// go through each byte of body
	// just locking the processing part concurrent
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)

		if cIndex >= 0 {
			// mutex.Lock()
			frequency[cIndex] += 1
			// mutex.Unlock()
		}
	}
	mutex.Unlock()
	fmt.Println("Completed: ", url, time.Now().Format("15:04:05"))

	// mutex.Unlock()
}
