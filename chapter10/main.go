package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

type FileLogger struct {
	File *os.File
}

func main() {
	// file, _ := os.Create("log.txt")
	// Process("Data to Process", FileLogger{File: file})
	// Process("Message to Log", ConsoleLogger{})
	MainHttp()
}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("cl_logger: %s\n", message)
}

func (f FileLogger) Log(message string) {
	fmt.Fprintf(f.File, "File: %s\n", message)
}

func Process(data string, logger Logger) {
	logger.Log("Started Processing...\n")
	logger.Log(data)
	logger.Log("Processing Finished\n")
}
