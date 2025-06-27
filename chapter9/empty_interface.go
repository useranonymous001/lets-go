package main

import "fmt"

// type Empty interface{}

/*
	The Empty Interface is the interface type that has no methods.
	Normally, Empty interface will be used in the literal form: "interface".
	All types satisfies Empty Interface
*/

func MainInterface() {
	Println("hello", "user")
	Println([]int{1, 2, 3, 5})
	Println(map[string]int{
		"hello": 12,
		"user":  01,
	})
}

// let's see an example of Println() function that accepts any kind of interface
// use case of empty interface
func Println(value ...interface{}) {
	fmt.Println(value...)
}

// Instead of using ...interface{} we can use it's alias any
// any => interface{}
