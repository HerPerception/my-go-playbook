package main

import "fmt"

type Person struct {
	Name       string
	Complexion string
	Age        int
	Gender     string
}

type PersonField interface {
	GetAge() int
	GetComplexion() string
	GetGender() string
	GetName() string
}

func (p Person) GetName() string {
	return p.Name
}
func (p Person) GetAge() int {
	return p.Age + 1
}
func (p Person) GetComplexion() string {
	return p.Complexion
}
func (p Person) GetGender() string {
	switch p.Gender {
	case "female":
		if p.Age <= 18 {
			return "Girl"
		} else if p.Age > 18 && p.Age < 30 {
			return "Lady"
		} else if p.Age >= 30 {
			return "Woman"
		}
	case "male":
		if p.Age <= 18 {
			return "Boy"
		} else if p.Age > 18 && p.Age < 30 {
			return "Young Man"
		} else if p.Age >= 30 {
			return "Man"
		}
	}
	return p.Gender
}
func NewAge(pf PersonField) {
	fmt.Printf("Happy Birthday %s, Today You Are %d years old.\n A wonderful %s, %s in complexion.", pf.GetName(), pf.GetAge(), pf.GetGender(), pf.GetComplexion())
}
func main() {
	var p Person
	p = Person{Name: "Daniella", Complexion: "dark", Age: 15, Gender: "female"}
	NewAge(p)
}
