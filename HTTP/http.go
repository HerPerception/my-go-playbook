package main

import "fmt"

// 1. Define the interface contract
type Speaker interface {
	Speak() string
}

// 2. Define a concrete type
type Dog struct {
	Name string
}

// 3. Implement the method (implicit satisfaction)
func (d Dog) Speak() string {
	return "Woof!"
}

// 4. Accept the interface as a parameter
func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	d := Dog{Name: "Buddy"}
	MakeItSpeak(d) // Works because Dog has a Speak() method
}
