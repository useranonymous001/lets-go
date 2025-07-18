package main

import (
	"fmt"
	"os"
	"sync"
)

/*

	Wait()
		- It unlocks the mutex so that other go routine can now access the shared variable
		- Automatically puts the goroutine to sleep until the condition is met or signal is received from another go routine

	Signal()
		- Now, the signal is sent by another go routine stating that i have made some changes
		checkout that you meet your required condition or not.

*/

// type Locker interface {
// 	Lock()
// 	Unlock()
// }

// type Cond
// 	func NewCond(l Locker) *Cond
// 	func (c *Cond) Broadcast()
// 	func (c *Cond) Signal()
// 	func (c *Cond) Wait()

func main() {
	// money := 100

	// mutex := sync.Mutex{}
	// cond := sync.NewCond(&mutex)
	// go Stingy(&money, cond)
	// go Spendy(&money, cond)
	// time.Sleep(2 * time.Second)

	// mutex.Lock()
	// fmt.Println("Bank Balance: ", money)
	// mutex.Unlock()

	MainMissingSignal()

}

// func Spendy(money *int, mutex *sync.Mutex) {
// 	for i := 0; i < 1000000; i++ {
// 		mutex.Lock()
// 		*money -= 10

// 		if *money < 0 {
// 			fmt.Println("Money is negative!")
// 			os.Exit(1)
// 		}

//			mutex.Unlock()
//		}
//		fmt.Println("Spendy Done")
//	}
//
// func Spendy(money *int, mutex *sync.Mutex) {
func Spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		// mutex.Lock()
		cond.L.Lock()

		// continuously check the balance if it is greater than 50 or not
		// unlocks the money variable and allow other go routien to modify it and then locks
		// again to check of it;s greater than 50 again
		for *money < 50 {
			// mutex.Unlock()
			// time.Sleep(10 * time.Millisecond) // it is not ideal idea
			// mutex.Lock()
			cond.Wait() // i am waiting, someone can now change the shared value
			// i wwant to meet my condition asap
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Negative Money")
			os.Exit(1)
		}
		// mutex.Unlock()
		cond.L.Unlock()
	}
	fmt.Println("Spendy Done")
}

// func Stingy(money *int, mutex *sync.Mutex) {
func Stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		// mutex.Lock()
		cond.L.Lock()
		*money += 10

		// signal that i have made some change to the shared variable
		// let's another go routine that would be waiting any kind of change in the shared variable
		// to meet its condition

		cond.Signal() // sends signal that check it now

		// mutex.Unlock()
		cond.L.Unlock()
	}
	fmt.Println("Stingy Done")
}
