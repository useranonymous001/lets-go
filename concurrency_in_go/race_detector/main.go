package main

import (
	"fmt"
	"time"
)

func main() {

	money := 100
	go stingy(&money)
	go spendy(&money)

	time.Sleep(2 * time.Second)
	fmt.Println("Money in the Bank: ", money)

}

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Spendy Done")

}
