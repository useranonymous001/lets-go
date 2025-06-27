/*
*** ARRAYS SLICES MAPS ***

	var arr [5]int        [0 0 0 0 0]
	data := [5]float64{}  [0 0 0 0 0]

*/

package main

import "fmt"

var x [5]int

func main() {
	x[4] = 100
	// fmt.Println(x)
	// average()
	// checkType()
	// practiceFunc()
	// Slices()
	// demoSlice()
	// Copy()
	// Maps()
	// advancedMaps()
	// Exercise()

	// from practice.go
	// array_practice()
	// slice_practice()
	// slice_optimization()
	map_practice()
}

// let's get some error first
func average() {
	size := 0
	fmt.Printf("Enter size of arr: ")
	fmt.Scanf("%v", &size)

	data := [5]float64{}
	// var data [size]int

	for i := 0; i < size; i++ {
		fmt.Printf("Enter data %v: ", i)
		fmt.Scanf("%f", &data[i])
	}

	// for i := 0; i < size; i++ {
	// 	fmt.Println(data[i])
	// }

	// sum of an array
	// total := 0.00
	var total float64 = 0
	// for i := 0; i < len(data); i++ {
	// 	total += data[i]
	// }

	for _, value := range data {
		// total += value
		fmt.Println(value)
	}

	fmt.Println("Total: ", total)
	fmt.Println("Average: ", total/float64(len(data)))

}

func checkType() {
	integers := 10
	floats := 10.10
	runeValue := 'c'
	strings := "hello world"

	fmt.Printf("10 %T\n", integers)
	fmt.Printf("10.10 %T\n", floats)
	fmt.Printf("c %T\n", runeValue)
	fmt.Printf("hello %T\n", strings)
	fmt.Printf("l %v\n", strings[2])
	fmt.Print("hello"[2])
}

func practiceFunc() {
	arr := [5]int{12, 23, 43, 54, 54}

	// when the iterable are not used like i
	/*
		A single underscore (_) is used to tell the compiler that
		we don’t need this (in this case, we don’t need the iterator variable).
	*/
	for _, value := range arr {
		fmt.Println(value)
	}

	// when the iterable is needed like i(index)
	for i, v := range arr {
		fmt.Println(i)
		fmt.Println(v)
	}

}

// Slices
func Slices() {

	// declaring slice

	// method 1: using make()
	nums := make([]int, 5) // creates a slice of 5 ints (all 0)
	// []int => means it's a slice of int
	fmt.Println(nums)

	numbers := make([]int, 5, 10)
	fmt.Println(numbers)
	// fmt.Println(len(numbers))  // 5
	// fmt.Println(cap(numbers))  // 10

	// method 2: From a literal
	names := []string{"rohan", "chris", "dallo"}
	fmt.Println(names)

	// method 3: slicing an existing array
	values := []int{1, 2, 3, 4, 5, 10}
	slice := values[1:3]
	fmt.Println("Actual values: ", values)
	fmt.Println("Slice values: ", slice)
	fmt.Println(len(values)) // 6
	fmt.Println(cap(values)) // 6

	// dynamicArr := [10]int{1, 2, 3, 4, 5}
	// dynamicArr = append(dynamicArr, 2, 7) // the array cannot be appended, it must be slice
	// fmt.Println(dynamicArr)

	staticArr := [10]int{1, 2, 3, 4, 5}
	dynSlice := staticArr[0:3]
	dynSlice = append(dynSlice, 10, 21)
	fmt.Println(staticArr)
	fmt.Println(dynSlice)

	fmt.Println(cap(staticArr))
	fmt.Println(cap(dynSlice))

}

func demoSlice() {
	// array := make([]int, 5, 10) // makes the size of an underlying array fixed which indirectly sets the capacity of the slice to be fixed as well

	// slice := []int{1, 2, 3, 4, 5} // this is actually a slice not array
	// slice = append(slice, 10, 11, 12, 13, 14)

	// slice := make([]int, 5, 10) // i can create slice this way as well
	// slice2 := []int{1, 2, 3, 4} // also this way too...
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := array[0:5]
	newSlice = append(newSlice, 100, 101, 102, 103, 104, 105)
	fmt.Println(newSlice)
	fmt.Println(array)
	fmt.Println("After re-allocation...")
	newSlice[0] = 10
	fmt.Println(newSlice)
	fmt.Println(array)

	// memory wastage, we are trying to create a new slice from a already created slice
	s1 := array[5:10]
	s2 := append(s1, 999)
	fmt.Println(s1, s2, array)
}

func Copy() {

	// slice := make([]int, 5, 10)
	// slice = append(slice, 10, 20) // [0 0 0 0 0 10 20] because i make() will create a slice with [0 0 0 0 0]

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	copySlice := make([]int, 5, 10)
	copy(copySlice, slice) // only copies upto the size of the destination slice
	fmt.Println(slice, copySlice)
}

func Maps() {
	// var x map[string]string // wrong way of definition of map
	// x := map[string]string{} // short hand for declaring a map

	x := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
	}

	if name, status := x["h"]; status {
		fmt.Println(name)
	} else {
		fmt.Println("Not Found")
	}

}

func advancedMaps() {

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
	}

	if el, status := elements["Be"]; status {
		fmt.Println(el["name"], el["state"])
	} else {
		fmt.Println("Not Found")
	}

}

func Exercise() {
	// x := [6]string{"a", "b", "c", "d", "e", "f"}
	// // creating a slice
	// slice := x[2:5]

	// fmt.Println(slice)

	x := []int{
		2,
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 1,
	}
	smallest := x[0]
	for _, v := range x {
		if v < smallest {
			smallest = v
		}
	}

	fmt.Println(smallest)
}
