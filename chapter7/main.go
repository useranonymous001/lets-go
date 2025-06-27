package main

import (
	"fmt"
	"reflect"
)

func main() {

	// defining custom data types using type keyword
	// type age int

	// userAge := age(10)
	// fmt.Println(userAge)
	// fmt.Printf("%T", userAge)

	// type Circle struct {
	// 	radius    float64
	// 	perimeter float64
	// }

	// // creating a local variable of type Circle
	// var c Circle
	// fmt.Println(c)

	// // short hand method
	// c2 := Circle{10.0, 20.0}
	// fmt.Println(c2)

	// // return pointer to the circle
	// c3 := new(Circle)
	// c3.perimeter = 10000
	// c3.radius = 1000
	// fmt.Println(*c3)

	// // another way of returning a pointer
	// c4 := &Circle{222, 333}
	// fmt.Println(c4, *c4)

	// struct_practice()
	// struct_embedding()
	// basic_struct()
	// AnonymousStruct()
	StructTags()
}

func struct_practice() {

	// an empty struct
	type Empty struct{}

	type Person struct {
		Name   string
		Age    int
		Salary float64
	}

	// initializing a struct
	var p1 Person
	p1.Name = "Rohan Khatri"
	p1.Age = 10
	p1.Salary = 100000.000
	fmt.Printf("%+v\n", p1)

	// using short hand values
	p2 := Person{} // initialzes an emmpty struct with each field to thier corresponding default value
	fmt.Printf("%+v\n", p2)
	p3 := Person{Name: "Rohan Khatri", Age: 10, Salary: 100.0100}
	fmt.Println(p3)
}

func struct_embedding() {

	type Person struct {
		Name  string
		Age   int
		Hobby []string
	}
	type Member struct {
		Person // this can be understood as, all the fields of the stuct Person is copied to the Member struct
		// that means we can access all the fields of the Person from the Member
		ID     int
		isHead bool
	}

	member := Member{
		ID:     1001,
		isHead: false,
		Person: Person{
			Name: "Rohan Khatri",
			Age:  20,
			Hobby: []string{
				"Football",
				"Pubg",
			},
		},
	}

	fmt.Println(member.Name)
}

type Animal struct {
	Name    string `required max:"100"`
	Origin  string
	Kingdom string
}

func StructTags() {

	t := reflect.TypeOf(Animal{})

	field, status := t.FieldByName("Name")
	fmt.Printf("%+v", field)
	fmt.Println(status)
	fmt.Println(field.Tag)
}
