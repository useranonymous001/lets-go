/*
Key Points to Remember:
	- defer function is executed, when the function is going to end
	- like just before the return statement

*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func deferPractice() {

	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("What the heck is happening...")
		}
	}()
	fmt.Println("Welcome to XSeed")
	fmt.Println("doing some processing....")

	panic("Woooohh, something went wrong, handle me")
}

// let's create our own defer and panic state, considering a real web-application

func Xseed() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, My first web go server"))
	})
	err := http.ListenAndServe(":8080", nil)

	defer func() {
		err := recover()
		if err != nil {
			panic("Fuck, the port already in use") // re-panicking
			// log.Fatal(err)
		}
	}()

	if err != nil {
		panic(err.Error())
	}
}

func Panicker() {

	fmt.Println("Something bad about to happen")

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("Ohh!! here we go, bad happened")
	fmt.Println("Done Panicking...")

}

func RealDeferUseCase() {
	// f, _ := os.Open("main.go")
	// defer f.Close()

	// buf := make([]byte, 1024)
	// if n, err := f.Read(buf); err == nil {
	// 	// fmt.Println(buf[:n])  byte of strings in buffer
	// 	fmt.Println(string(buf[:n]))
	// }

	res, err := http.Get("https://www.github.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}
	// close the res.Body before the function ends
	defer res.Body.Close()

	robots, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

}
