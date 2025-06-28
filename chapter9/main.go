package main

import "fmt"

/*
Interface in GO
	Interface in go is like a blueprint that specifies the behaviour of an object.

	Object => Instance of its type

*/

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Stringer interface {
	String() string
}

/*
If any type satisfy this interface - that is define these two methods
which returns float64 - then, we can say that type is implementing
this interface.

*/

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

/*
In Go implementing an interface happens implicitly.
So, no need to explicitly declare a type is implementing a particular interface


As you can see above, the Rectangle type has two methods named Area() and Rectangle() which
returns float64. So, we can say that the type Rectangle is implement`ing "Geometry Interface".
*/

func main() {
	// r := Rectangle{Height: 10.10, Width: 34.12}
	// c := Circle{Radius: 8.9}
	// fmt.Println(r.Area(), r.Perimeter())
	// fmt.Println(c.Area(), c.Perimeter())
	// fmt.Println("Dimension of", r.String())
	// fmt.Println(r.Area(), r.Perimeter())

	// Measure(r)
	// Measure(c)

	// MainInterface()
	// PracticeMain()
	// TypeAssertionMain()
	MainInterfaceComposition()

}

// methods on type Rectangle  (implements Geometry Interface)
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height * r.Width)

}

// methods on type Circle (implements Geometry Interface)
func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// methods on type Rectangle (implements Stringer Interface)
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectange: %v x %v", r.Height, r.Width)
}

/*
So, if anywhere the Geometry interface type is expected, you can use any of
these implementations
*/

func Measure(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}

/*
When you call the above function, you can pass the argument as
an "object of Geometry interface type". Since both Rectangle and
Circle satisfy that interface, you can use either one of them.
*/

// Type can also implement multiple Interface
// Type with multiple interfaces
