package main

import "fmt"

func array_practice() {
	fmt.Println("Practicing arrays")

	arr1 := [3]int{}
	// arr2 := [3]int{}

	arr1[1] = 10
	fmt.Println(arr1)

	// important syntax
	// there's a way to declare a array literal without specifying the length
	arr3 := [...]int{12, 31, 123, 100, 342}
	fmt.Println(arr3, len(arr3))

	for index, value := range arr3 {
		fmt.Println(index, value)
	}
}

func slice_practice() {
	// slices need to be initilized with size more than 0
	slice := []int{} // throws a panic error, index out of ranges
	// slice[0] = 100   // "

	var slice2 []string

	fmt.Printf("%#v, %#v\n", slice, slice == nil)   // []int{}, false
	fmt.Printf("%#v, %#v\n", slice2, slice2 == nil) // []string(nil), true

	// but how can we initilize a slice with some value but still empty
	// we can make use of built-in function called make()

	slice3 := make([]string, 0, 5)   // ["" "" ""]
	slice3 = append(slice3, "rohan") //["" "" "" "rohan"]
	// fmt.Println(len(slice3), cap(slice3)) // 4 5

	// appending one slice to another slice
	slice4 := make([]string, 0, 10)
	slice4 = append(slice4, slice3...)
	// fmt.Println(slice4)

	// slicing from an already existing array
	users := [...]string{
		"rohan",
		"bandana",
		"amrit",
		"auraj",
		"ningma",
		"binayak",
		"alwin",
		"samir",
		"anuja",
		"ashika",
		"kuber",
		"nischal",
		"arjun",
	}

	fruits := []string{
		"mango",
		"banana",
		"apple",
		"grapes",
		"papaya",
		"coconut",
		"litchi",
		"pearls",
	}

	// users = append(users, "newusre") // can't do this to an array
	fmt.Println(users)
	sectionA := users[:7]
	fmt.Println("Before appending", len(sectionA), cap(sectionA))
	sectionA = append(sectionA, fruits...)
	fmt.Println(sectionA)
	fmt.Println("After appending", len(sectionA), cap(sectionA))
	fmt.Println(users, len(users), cap(users))

	/*
		When the slice length exceeds the capacity of it's underlying array then
		it re-allocates the new array with double of it's current size, leaving
		already existing array as it is and creating a new slice with double of its capacity.
	*/

	// copying the slice to another slice
	// newSlice := []string{}//
	newSlice := make([]string, len(sectionA), cap(sectionA))
	copy(newSlice, sectionA) // cannot copy to the slice that have size 0
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
}

func slice_optimization() {

	// not optimized
	// slice := make([]int, 0)
	// for i := 0; i < 90000000; i++ {
	// 	slice = append(slice, i)
	// }
	// fmt.Printf("%v\n", len(slice))

	/// little optimization

	slice := make([]int, 0, 90000000)
	for i := 0; i < 90000000; i++ {
		slice = append(slice, i)
	}
	fmt.Println(len(slice))

}
