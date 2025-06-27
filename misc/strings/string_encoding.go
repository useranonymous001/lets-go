package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "🐻"

	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(len(str))

	fmt.Printf("Unicode: U+% 04X\n", str)
	fmt.Printf("Bytes in UTF-8: % X\n", []byte(str))

	// for _, v := range str {
	// 	fmt.Printf("%v ", v)
	// 	fmt.Println("")
	// }

	initialStatemen()

	fmt.Println(founders("linux"))
	fmt.Println(founders("acer"))

}

func getName() (string, bool) {
	return "Rohan", true
}

func initialStatemen() {

	if name, status := getName(); status {
		fmt.Println(name, status)
	}
}

func founders(os string) string {
	var creator string

	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "window":
		creator = "Bill Gates"
	case "acer":
		fallthrough
	case "hp":
		fallthrough
	case "dell":
		creator = "Rohan Khatri"
	default:
		creator = "Unknown"
	}

	return creator
}
