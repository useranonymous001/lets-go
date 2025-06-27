/*
*** CONTROL STRUCTURE ***


 "FOR" STATEMENT
	- just like for loops in other programming
	- for (condition) {
		// program statements
	}

	*** other way ***
	for initialization ; condition ; increment/decrement {

	}

	*** SWITCH CASES ***
	switch expression {
		case 1:
			so something...
		default :
			do something..

	}

*/

package main

import "fmt"

func main() {
	fmt.Println("We are learning Control Structures in GO")
	// forStatement()
	// switchStatement()
	// divisibleByThree()
	// fizzBuzz()
	shortCut()
}

func forStatement() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println("even")
		} else if i%2 != 0 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Not a number")
		}
	}
}

func switchStatement() {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Un recognizable")
		}
	}
}

func divisibleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		}
	}
}

func shortCut() {

	// length := len("rohan@gmail.com")
	// if length < 5 {
	// 	fmt.Println("Not a valid email")
	// } else {
	// 	fmt.Println("Valid email")
	// }

	if length := len("rohan@gmail.com"); length > 5 {
		fmt.Println("Valid email")
	}
}
