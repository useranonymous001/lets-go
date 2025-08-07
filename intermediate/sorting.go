package main

import (
	"cmp"
	"fmt"
	"slices"
)

// sort functions are generic means they acccepts any kind of ordered types like int float int64
func MainPackage() {
	sts := []string{"a", "b", "c"}
	slices.Sort(sts)
	fmt.Println(sts)
	ints := []int{11, 222, 3, 42, 7}
	slices.Sort(ints)
	fmt.Println(ints)

	SortingFunc()
}

func SortingFunc() {

	type Employee struct {
		id     int
		salary int
	}

	emplpoyees := []Employee{
		Employee{id: 101, salary: 1010101},
		Employee{id: 102, salary: 50001},
		Employee{id: 103, salary: 90121},
		Employee{id: 104, salary: 49200},
	}

	slices.SortFunc(emplpoyees, func(a, b Employee) int {
		return cmp.Compare(a.salary, b.salary)
	})

	fmt.Println(emplpoyees)
}
