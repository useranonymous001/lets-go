package main

import (
	"fmt"
	"regexp"
)

func MainRegex() {
	var print = fmt.Println
	matched, _ := regexp.MatchString("ro([a-z]+)n", "rohan")
	print(matched)
}
