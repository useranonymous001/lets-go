package main

import (
	"fmt"
	"os"
	"searchEngine/search"
)

func main() {
	query := os.Args[1]
	engine := search.InitializeEngine("./search/data")
	result := engine.Search(query)
	fmt.Println("Total Result Found: ", len(result))
	for key, value := range result {
		fmt.Printf("%s: %s\n\n", key, value)
	}
}
