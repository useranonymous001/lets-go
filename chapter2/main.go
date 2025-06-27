/*
*** TYPES ***

	1) Go is statically typed language, meaning its types cannot be changed

	2) Types
		- Numbers
			- Integers
			- Floating Point Numbers
		- String
		- Boolean

*/

package main

import (
	"fmt"
)

func main() {
	/*
		myAge := 2131232234234223423 // by default its int
		yourAge := -11231231231241242348327428347246242834762384264283491263471298364712934813647129834163471298341364712893127634836471283412734812634912316934182364712387631544387634573453475.12312123123123123421312312341231423847247623842472534
		fmt.Println(myAge)
		fmt.Println(math.Abs(yourAge))
		fmt.Println("1 + 1 = ", 1.1+1)
	*/

	newStr := "Hello, this is captain joe speaking from go"
	backticks := ` and this is backticks from go`

	fmt.Println(backticks)
	fmt.Println(newStr)
	fmt.Println("Length of newStr: ", len(newStr))
	// concatenation
	fmt.Println(newStr + backticks)
	fmt.Println(newStr[1])

	fmt.Println(true || false)
	fmt.Println(!false)

	num1 := 32132
	num2 := 42452

	fmt.Println(num1 * num2)

}
