/*
*** METHODS IN GO FUNCTIONS ***
	Functions in GO can also be used as types
		they can be used as the arguments, parameters  or any kind of values
*/

package main

import "fmt"

func main() {

	// anonoymousFunction()
	// funcAsVariable()
	// BasicMethods()
	// MethodPractice()
	MethodOnStruct()
}

func anonoymousFunction() {
	// it's a good practice to pass the argument to the anon func if we are sometime working with anonymous function
	for i := 0; i < 5; i++ {
		func(i int) { // good practice to actually pass the parameters to the anon functions as well
			fmt.Println(i)
		}(i)
	}
}

func funcAsVariable() {
	var divide func(float64, float64) (float64, error)

	divide = func(v1, v2 float64) (float64, error) {
		if v2 == 0 {
			return 0.0, fmt.Errorf("Divisor cannot be 0")
		}
		result := v1 / v2
		return result, nil
	}
	res, err := divide(20.121, 8.0)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
