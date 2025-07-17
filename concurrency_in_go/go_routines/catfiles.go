/*

Write a program similar to the one in listing 2.3 that accepts a list of text file-
names as arguments. For each filename, the program should spawn a new
goroutine that will output the contents of that file to the console. You can use
the time.Sleep() function to wait for the child goroutines to complete (until
you know how to do this better). Call the program catfiles.go. Here’s how you
can execute this Go program:

go run catfiles.go txtfile1 txtfile2 txtfile3

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func catFiles(files ...string) {

	for _, file := range files {
		go func(file string) {

			f, err := os.Open(file)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			buf := make([]byte, 1024)
			for {
				n, err := f.Read(buf)
				if err != nil {
					if err == io.EOF {
						break // end of file reached
					}
					fmt.Println("Error reading file:", err)
					return
				}
				fmt.Println(string(buf[:n])) // convert the buffer data to string

			}

		}(file)
	}

}

func grepFiles(files ...string) {

	for _, file := range files {

		go func(file string) {

			fs, err := os.Open(file)
			if err != nil {
				fmt.Println("Eror opening a file")
				return
			}

			scanner := bufio.NewScanner(fs)
			for scanner.Scan() {
				line := scanner.Text()

				if strings.Contains(line, "bubbles") {
					fmt.Println("Found at file: ", file)
					return
				}
			}

		}(file)

	}

}
