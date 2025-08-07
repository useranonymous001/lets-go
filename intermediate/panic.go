package main

import (
	"fmt"
	"os"
)

func MainPanic() {
	// panic("what the hell")
	// _, err := os.Create("/tmp/err")
	// if err != nil {
	// 	panic(err)
	// }

	MainDefer()
}

func MainDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("baal baal bachgaya", r)
			return
		}
	}()
	f := CreateFile("/temp/temp.txt")
	defer CloseFile(f)
	WriteFile(f)
}

func CreateFile(path string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	return f
}

func WriteFile(f *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(f, "hello honey bunny")
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
