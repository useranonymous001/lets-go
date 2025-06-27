/*
*** FUNCTIONS IN GO ***


CLOSURES:

Definition and Explanation 123

A closure is a function with an extended scope that encompasses variables referenced in the body of the function that are not global variables or local variables of the function. This allows the function to access and modify variables from outside its scope, making it a powerful tool for programming. Closures are sometimes confused with anonymous functions, but a closure is a function that can access nonglobal variables defined outside of its body.

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(f1())
	// fmt.Println(multipleValues())
	// fmt.Println(nameReturnType())

	// if returning multiple values, accept in this way
	// if x, status := multipleValues(); status {
	// 	fmt.Println(x, status)
	// }

	// variadic parameters to the function
	// variabicFunc('A', '2', 3, 4, 5)
	// x := []int{1, 2, 3, 4, 5}
	// variabicFunc(x...)

	// Closure
	// addClosure()
	// closure := makeClosure()
	// fmt.Println(closure())
	// fmt.Println(closure())
	// fmt.Println(closure())

	// defer variabicFunc(x...) // it will be called at the very end
	// addClosure()
	// fmt.Println(nameReturnType())
	// deferUseCase()
	// panicUsage()
	// constUsage()

	// practicalUseCase()
	// Pointers()

	// from def_pan_recover.go
	// deferPractice()
	// Xseed()
	// fmt.Println("Hello There...")
	// Panicker()
	// fmt.Println("End!!")
	// RealDeferUseCase()
	DefaultPointer()
}

// stack implementation
func f1() int {
	return f2()

}

func f2() int {
	fmt.Println("func 2")
	return 1
}

func multipleValues() (int, bool) {
	return 5, false
}

func nameReturnType() (num int) {
	num = 1
	return
}

func variabicFunc(args ...int) {
	fmt.Println(args)
}

func passAnArray(array []string) []string {
	arrayOfString := []string{"helo", "world"}
	return arrayOfString
}

func addClosure() {

	localVar := 10
	add := func(x, y int) int {
		localVar++
		return x + y + localVar
	}

	fmt.Println(add(2, 3))
	fmt.Println(localVar)
	fmt.Println(add(2, 3))
	fmt.Println(localVar)

}

func makeClosure() func() uint {
	state := uint(0) // can access this state using the inner function easily
	return func() (ret uint) {
		ret = state
		state += 2
		return
	}
}

// defer panic and recover

// defer:  it is an special statement in a function that is scheduled to be called after the function completes. Basically, it moves the call to the function with (defer) statement at the very end of the program/main function

// use cases: often used when we need to free spaces/resources after used
// closing file after opening and using it
// Deferred functions are run even if a runtime panic occurs

func deferUseCase() {
	f, _ := os.Open("main.go")
	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]))
}

func panicUsage() {
	// panic("what the fuck??") // automatically stops the execution of the program
	// str := recover()
	// fmt.Println(str)

	// so use with defer

	defer func() {
		str := recover()
		fmt.Println(str)
		fmt.Println("Gotch you babyyy ")
	}()
	panic("Oh no fucking godd, panic panic, someone catch me catch me...")

	/*
		Explanation why, it works with defer

			panic("no nn")
			str := recover() // This line is never reached
			As soon as panic("no nn") is called, the program starts unwinding the stack.

			It looks for any deferred functions to run before it crashes.

			The line str := recover() is never reached, because the panic has already begun.

			Even if it were reached, recover() won’t recover unless it’s inside a deferred function.
	*/

}

func practicalUseCase() {

	// error handling
	defer func() {

		errMessage := recover()
		// fmt.Printf("%T\n", errMessage)
		// fmt.Printf("%T\n", errMessage.(string))

		if errMessage != nil {
			if msg, ok := errMessage.(string); ok && len(msg) > 0 {
				// fmt.Println(ok)
				fmt.Println(msg)
			} else {
				fmt.Println(errMessage)
			}
		}

	}()

	costPerMessage := 2
	accountBalance := 88
	messageCount := 0
	totalCost := 0

	fmt.Printf("Enter number of message to send: ")
	fmt.Scanln(&messageCount)

	totalCost = messageCount * costPerMessage

	if totalCost > accountBalance {
		panic("Brooo, recharge your balance first")
	} else {
		accountBalance -= totalCost
		fmt.Println("Sending message...")
		fmt.Println("Available Balance: ", accountBalance)
	}

}

func Pointers() {
	// a, b := 10.0, 20.0
	// fmt.Println(p_add(&a, &b))
	// fmt.Println(a, b)

	// another way of getting a pointer
	xPtr := new(int)   // xPtr is a pointer a new memory location that stores integers values
	fmt.Println(xPtr)  // 0xc000098040
	fmt.Println(*xPtr) // 0
	one(xPtr)
	fmt.Println(*xPtr) // 1000
}

// helper functions for pointers demonstrations
func p_add(a, b *float64) float64 {
	*a += 10
	*b -= 3

	fmt.Println(a, b)

	return *a + *b
}

func one(xPtr *int) {
	*xPtr = 1000
}

func demo() {

	add := func() {
		fmt.Println("hello world")
	}

	defer func() int {
		return 0
	}()
	add()
	// increment()
}
