/*
	Mutexes (Mutual Exclusions)
	(sometimes called locks)

	Mutexes protects the critical section of our code so that only one go routine can access the critical section at a time.

	Varients of Mutexes

	- Reader-Writer Mutexes
	 	It gives the performance optimizations in situations where we need to block concurrency only when modifying the shared resources.

		That means allowing multiple readers to the shared resources but only one writer at a time.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	// money := 100
	// mutex := sync.Mutex{} // initial value is Unlocked
	// go spendy(&money, &mutex)
	// go stingy(&money, &mutex)
	// time.Sleep(time.Second * 2)
	// fmt.Println("Bank balance: ", money)

	// MainMutexes()
	MainRwMutex()

}

func spendy(money *int, mutex *sync.Mutex) {

	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}

	fmt.Println("Spendy Done")
}

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Stingy Done")
}
