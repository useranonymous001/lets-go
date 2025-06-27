package main

import "fmt"

// Upper case in the field names and and struct name for export
// use Lower case for the import

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func basic_struct() {

	aDoctor := Doctor{
		Number:    1,
		ActorName: "John Smith",
		Companions: []string{
			"chris",
			"joe",
			"amilia",
		},
	}
	fmt.Println(aDoctor.Companions[1])

}

func AnonymousStruct() {
	// this is something like when you just need struct for small amount of time.

	aDoctor := struct{ name string }{name: "Rohan Khatri"}
	// but unlike maps and slices, when you pass struct to another variable or any function,
	// you do not pass the reference of the struct but the copy of the actual struct
	anotherDoctor := aDoctor
	anotherDoctor.name = "John Hamilton"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	// you need to pass with the pointer
	// something like this way
	newDoctor := &aDoctor
	newDoctor.name = "Warren Buffet"
	fmt.Println(newDoctor)
	fmt.Println(aDoctor)

}
