/*

Variables

	Don’t let the best be the enemy of the good.

	Go is lexically scoped using blocks.

*/

package main

import "fmt"

var anyvar int // globally available

func main() {

	stringVar := "helo world"
	var name string = "hello world"
	// var rollNo = 12

	fmt.Println(stringVar == name)
	fmt.Println(name)
	// fmt.Println(anyvar)
	// newFunc()
	// Example()
	// tempConverter()
	distanceConverter()
}

func Example() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func newFunc() {
	const PI float64 = 3.1415

	// shorthand method for declaring variables

	var (
		rol = 123
		mol = 456
		gol = 234
	)

	const (
		pi    = 3.1415
		messi = 10
	)

	fmt.Println(PI)
	fmt.Println(anyvar)
	fmt.Println(rol)
	fmt.Println(mol)
	fmt.Println(gol)
	fmt.Println(pi)
	fmt.Println(messi)
}

func tempConverter() {

	var Fahrenheit float64
	fmt.Printf("Enter temp in F: ")
	fmt.Scanf("%f", &Fahrenheit)
	fmt.Println(Fahrenheit)
	Celcius := ((Fahrenheit - 32) * 5 / 9)

	fmt.Printf("%.2fF = %.2fC\n", Fahrenheit, Celcius)
}

func distanceConverter() {
	var ft float64
	fmt.Print("Enter distance in feet: ")
	fmt.Scanf("%f", &ft)

	meter := ft * 0.3048

	fmt.Printf("%.2fFt = %.2fm", ft, meter)
}
