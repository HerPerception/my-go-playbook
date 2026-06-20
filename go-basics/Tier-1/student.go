package main

import "fmt"

type Student struct {
	Name  string
	Grade float64
}

func HasPassed(student Student) bool {
	if student.Grade >= 50.0 {
		return true
	}
	return false
}
func Real() {
	st1 := Student{
		Name:  "Cilia",
		Grade: 49.5,
	}
	st2 := Student{
		Name:  "Datty",
		Grade: 70.1,
	}
	fmt.Println(HasPassed(st1))
	fmt.Println(HasPassed(st2))
}
