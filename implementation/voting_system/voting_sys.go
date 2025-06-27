package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("*** Welcome to lazi's voting app *** ")
	fmt.Println(" ")
	// votingSystem()
	test()
}

func votingSystem() {

	candidates := []string{"Prachandey", "KP Budo", "Balen", "Sherey", "Don't Vote"}
	votes := map[string]int{}
	option := 2

	for i, candidate := range candidates {
		fmt.Printf("%d: %s\n", i+1, candidate)
	}
	fmt.Println("")
	fmt.Print("Enter the candidate number: ")
	fmt.Scanln(&option)

	if option == 5 {
		fmt.Println("Don't want to vote")
		os.Exit(1)
	}

	fmt.Printf("\nVoting %s....", candidates[option-1])
	chosedCan := candidates[option-1]
	votes[chosedCan] = votes[chosedCan] + 1

	fmt.Println("")
	fmt.Println("\n*** Result ***")
	for c, vote := range votes {
		fmt.Printf("%s : %d", c, vote)
	}

	fmt.Println("")

}

func test() {

	slice1 := []int{1, 2, 3, 4}
	// slice2 := []int{5, 6, 7, 8}
	slice2 := make([]int, 5)
	// slice2 = append(slice2, slice1...)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
