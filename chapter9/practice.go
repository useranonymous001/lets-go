package main

import "fmt"

type Temperature struct {
	Value    float64
	Location string
}

type Animal struct {
	Name  string
	Count int
}

func PracticeMain() {
	temp := Temperature{Value: 120.12, Location: "Jhapa"}
	bird := Animal{Name: "Humming Bird", Count: 12}
	// fmt.Println(temp.String())
	// fmt.Println(temp)
	cityTemperature(&temp)
	cityTemperature(bird)
}

func (b Animal) String() string {
	return fmt.Sprintf("Name: %s, Count: %d", b.Name, b.Count)
}

func (t *Temperature) String() string {
	(*t).Value = 21.2
	(*t).Location = "Kathmandu"
	return fmt.Sprintf("Temp: %.3f, Loc: %s", t.Value, t.Location)
}

func cityTemperature(v Stringer) {
	fmt.Println(v.String())
}
