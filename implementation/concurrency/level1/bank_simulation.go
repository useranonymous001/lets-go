/*
2. Bank Account Simulation
Simulate a bank account where multiple goroutines deposit and withdraw money.

Protect shared balance with mutex.

Ensure that a withdrawal doesn't happen if balance is insufficient.

✅ Use condition variable to wait for sufficient balance before withdrawing
✅ Prevent race conditions on balance
*/

package main

import (
	"fmt"
	"sync"
)

func MainBankSimulation() {

	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	balance := 0

	wg := sync.WaitGroup{}

	go func() {
		withdraw(&balance, 10000, cond)
		fmt.Println("Withdrawn successfull of 10000")
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			deposit(&balance, 1000, cond)
		}()

	}

	wg.Wait()

	cond.L.Lock()
	fmt.Println("Available balance: ", balance)
	cond.L.Unlock()
}

func withdraw(balance *int, amount int, cond *sync.Cond) {
	cond.L.Lock()
	for *balance < amount {
		cond.Wait()
	}
	*balance -= amount
	cond.L.Unlock()
}

func deposit(balance *int, amount int, cond *sync.Cond) {
	cond.L.Lock()
	*balance += amount
	cond.Signal()
	cond.L.Unlock()
}
