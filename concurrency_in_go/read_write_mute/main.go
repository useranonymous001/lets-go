package main

import (
	"fmt"
	"time"
)

// use case of ReaderWriterMutex with reader preferred mutex
// func main() {

// 	rwmutex := ReaderWriterMutex{}

// 	for i := 0; i < 2; i++ {
// 		go func() {
// 			for {
// 				rwmutex.ReadLock() // holds the reader lock
// 				time.Sleep(time.Second)
// 				fmt.Println("Read Done")
// 				rwmutex.ReadUnlock()
// 			}
// 		}()
// 	}

// 	time.Sleep(time.Second * 1)
// 	// after the reading is done, try to aquire the writer lock
// 	rwmutex.WriteLock()
// 	fmt.Println("write finished")

// }

func main() {
	rwmutex := ReadWriteMutex{}
	wmutex := rwmutex.NewReadWriteMutex()

	for i := 0; i < 2; i++ {
		go func() {
			for {
				wmutex.ReadLock()
				time.Sleep(time.Second)
				fmt.Println("Read Done")
				wmutex.ReadUnlock()
			}
		}()
	}

	time.Sleep(time.Second)
	wmutex.WriteLock()
	fmt.Println("write done")

}
