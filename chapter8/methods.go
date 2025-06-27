package main

import (
	"fmt"
	"os"
	"strconv"
)

/*

METHODS:
	- Methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct


func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
*/

type ByteSlice []byte

type Number int

func BasicMethods() {
	// s := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(s)
	// slice := ByteSlice{32, 43, 34, 12}
	// data, _ := slice.Append([]byte(s))
	// fmt.Println(data)
	// fmt.Println(cap(slice))

	// example from the go-lang book
	i := os.Args[1]
	n, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println("Not a number")
		os.Exit(1)
	}
	num := Number(n)
	fmt.Println(num.Even())

}

func (p *ByteSlice) Append(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		newSlice := make([]byte, 2*(len(data)+l))
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(slice), nil
}

// another practice on methods use case
type Profile struct {
	Name     string
	Follower int
	Bio      string
}

func MethodPractice() {
	var user Profile
	user.Name = "Rohan Khatri"
	user.Follower = 10
	user.Bio = "I am a gentle minimalistic guy"
	user.View() // same as: (&user).View()
	fmt.Println(user)
}

func (p *Profile) View() {
	(*p).Name = "Dalley Mishra" // same as p.Name = "Dalley Mishra"
	fmt.Printf("%+v\n", *p)
}

func (num Number) Even() bool {

	if num%2 == 0 {
		return true
	}

	return false

}

type Rectangle struct {
	Height float64
	Width  float64
}

func MethodOnStruct() {

	var rect Rectangle
	rect.Height = 10.10
	rect.Width = 20.12

	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Println(area, perimeter)

}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Height * rect.Width)
}
