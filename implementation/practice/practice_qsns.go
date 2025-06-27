package main

import "fmt"

func main() {
	fmt.Println("*** Practice time ***")
	fmt.Println("")
	// reverseArray()
	// frequencyCounter()
	// MaxMin()
	// checkPalindrome()
}

func reverseArray() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := append(array[:])
	fmt.Println(array)
	fmt.Println(slice)

	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}

	// for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
	// 	slice[i], slice[j] = slice[j], slice[i]
	// }

	fmt.Println(array)
	fmt.Println(slice)
}

func frequencyCounter() {
	input := ""
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)

	frequency := map[rune]int{}
	// create a map of the unique element and keep adding the repition of that character
	// for _, char := range input {
	// 	fmt.Printf("%c\n", char)
	// }

	for _, char := range input {
		frequency[char] = frequency[char] + 1
	}

	for char, count := range frequency {
		fmt.Printf("%c : %d\n", char, count)
	}
}

func MaxMin() {
	slice := []int{2, 3, 42, 1, 523, 23, 2576, 779, 87}

	largest, smallest := slice[0], slice[0]

	for _, v := range slice {
		if smallest > v {
			smallest = v
		} else if largest < v {
			largest = v
		}
	}

	fmt.Println(smallest)
	fmt.Println(largest)
}

func checkPalindrome() {
	inputStr := ""
	fmt.Print("Enter a string: ")
	fmt.Scanln(&inputStr)

	lastIndex := len(inputStr) - 1
	for i := 0; i < len(inputStr)/2 && i < lastIndex-i; i++ {
		if inputStr[i] != inputStr[lastIndex-i] {
			fmt.Println("Not Palindrome")
			return
		}
	}
	fmt.Printf("%s is Palindrome\n", inputStr)

}
