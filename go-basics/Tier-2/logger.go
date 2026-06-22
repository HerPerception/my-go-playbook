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
	Filename string
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	fmt.Fprintln(file, message)
}

func RunApp(l Logger) {
	l.Log("app started")
}

func TaskTwo() {
	h := ConsoleLogger{}
	file := FileLogger{Filename: "file.txt"}
	RunApp(h)
	RunApp(file)
}
