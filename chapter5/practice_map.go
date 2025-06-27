package main

import "fmt"

func map_practice() {
	fmt.Println("Practicing Maps...")

	// creating a map with valid and invalid key types

	// valid key type: string, int, bool, and arrays too
	// invalid key type: slice

	// m := map[[]int]string // throws an error
	// m := map[[2]int]string{}

	statePopulation := map[string]int{
		"Jhapa":      101023,
		"Morang":     123012,
		"Sunsari":    328572,
		"Kathmandu":  972934,
		"Biratnagar": 713641,
		"Bhaktapur":  165231,
		"Lalitpur":   67129,
		"Karnali":    862344,
		"Pokhara":    716345,
		"Illam":      91248,
		"Solukhumbu": 17624,
	}

	// it always maps the string in alphabetical order
	fmt.Println(statePopulation)
	statePopulation["Rukum"] = 51263 // this is actually returning two values
	// deleteing element from the map
	delete(statePopulation, "Lalitpur")
	fmt.Println(statePopulation)

	// pop, status := statePopulation["Ktm"]
	// fmt.Println(pop, status)

	// whenever you pass the map through any kind of functins, you are actually passing the reference to that map

	sp := statePopulation
	delete(sp, "Bhaktapur")
	fmt.Println(statePopulation)

	test_map(statePopulation)
	fmt.Println(statePopulation)

}

func test_map(sp map[string]int) {
	fmt.Println(sp)
	delete(sp, "Biratnagar")
	fmt.Println("inside the test: ", sp)
}
