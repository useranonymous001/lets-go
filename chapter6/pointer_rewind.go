package main

import "fmt"

func DefaultPointer() {

	// // slices
	// slice := []int{1, 23, 2, 4}
	// newSlice := slice
	// fmt.Println(slice, newSlice)
	// // newSlice[0] = 100
	// newSlice = append(newSlice, 100)
	// newSlice[0] = 222
	// fmt.Println(slice, newSlice)
	// fmt.Println(cap(slice), cap(newSlice))

	// maps
	// product := map[string]int{"Phone": 10000, "Laptop": 201923}
	// priceList := product
	// fmt.Println(product, priceList)
	// priceList["Phone"] = 99999
	// fmt.Println(product, priceList)
	// priceList["Watch"] = 76543
	// fmt.Println(product, priceList)

	// but this doesn't apply to arrays and structs

	p := new(Person)
	p.foo = "ehoqwe"
	fmt.Println(p.foo)

	p2 := p
	p2.foo = "hello world"
	fmt.Println(p2.foo)

	personPtr := &Person{}
	personPtr.foo = "ejalkfuhea"
	fmt.Println(personPtr.foo)

	// getting the pointer of the result
	res := fuckBoy()
	fmt.Println(res, &res)

}

type Person struct {
	foo string
}

func fuckBoy() *int {
	result := 1010101
	return &result
}
