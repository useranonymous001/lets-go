package main

import "fmt"

/*
Type assertion gives the underlying concrete value of an interface
type.


In the previous  example of Rectangle struct that implements Geometry interface,
if you want to print the width and and height
from the Measure function, you can use type assertions.

*/

func TypeAssertionMain() {
	r := Rectangle{Height: 83, Width: 3}
	c := Circle{Radius: 212}

	Measurement(r)
	Measurement(c)
}

func Measurement(g Geometry) { // accepts an object that implements Geometry Interface
	// rect := g.(Rectangle) // it throws panic if assertion failed
	// // alternative good way of handing values
	// if rect2, ok := g.(Rectangle); ok {
	// 	fmt.Println(rect2.Height)
	// 	fmt.Println(rect2.Width)
	// }
	// fmt.Println(rect.Height)
	// fmt.Println(rect.Width)

	switch v := g.(type) {
	case Rectangle:
		{
			fmt.Println("Height: ", v.Height)
		}
	case Circle:
		{
			fmt.Println("Radius: ", v.Radius)
		}

	default:
		fmt.Println("helw")
	}

}
