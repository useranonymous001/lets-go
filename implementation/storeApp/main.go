package main

import (
	"fmt"
	"log"
	"storeApp/storage"
)

func main() {
	// store := storage.NewFileStorage("data.txt")
	store := storage.NewMemoryStorage()

	err := store.Save("name", "anonymous001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Created And Read Successfully")
	str, lErr := store.Load("location")

	if err != nil {
		log.Fatal(lErr)
	}
	fmt.Println("str: ", str)
}
