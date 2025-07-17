package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	go doWork(i)
	// }
	// time.Sleep(2 * time.Second)
	// checkNumOfCPU()

	// go sayHello() // assigns the other goroutine a chance to run
	// call the scheduler directly in our main goroutine
	// it is not sure over which goroutine will select to execute
	// It may or may not select SayHello() function
	//  or it might continue the execution of the goroutine that called the scheduler.
	// runtime.Gosched()
	// fmt.Println("Finished")

	// dir_name := os.Args[1]
	// fmt.Println(dir_name)
	dir_name, _ := os.Getwd()
	dirEntries, err := os.ReadDir(dir_name)
	fmt.Println(dirEntries)
	if err != nil {
		fmt.Println("Dir not found")
		return
	}
	files := []string{}

	for _, v := range dirEntries {
		info, _ := v.Info()     // returns the fs.FileInfo for the file or subdirectory
		filename := info.Name() // returns base name of the file
		files = append(files, filename)
	}
	// catFiles(files...)
	grepFiles(files...)
	// runtime.Gosched()
	time.Sleep(2 * time.Second)
}

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05")) // Format returns a textual representation of the time value formatted according to the layout defined by the argument.
	time.Sleep(1 * time.Second)                                              // sleep for 1 second
	fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func checkNumOfCPU() {
	// NumCPU returns the number of logical CPUs usable by the current process.
	// The set of available CPUs is checked by querying the operating system
	// at process startup. Changes to operating system
	// CPU allocation after process startup are not reflected.
	fmt.Println("Number of CPUs: ", runtime.NumCPU()) // returns the number of logical cpus usable by the process

	// with n < 1 returns the current value without altering it.
	// GOMAXPROCS sets the maximum number of CPUs that can be
	// executing simultaneously and returns the previous setting.

	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}

func sayHello() {
	fmt.Println("hello boys")
}

// I have to find bubbles, bubble
// I have to find buble in the file main.go
