package main

import "fmt"

// A struct is a container that holds related data of possibly different types and groups them into a single unit or custom type.
type Animals struct {
	Name     string
	Sound    string
	Domestic bool
}

func main() {
	animals := []Animals{
		{Name: "Dog", Sound: "Bark", Domestic: true},
		{Name: "Cat", Sound: "Meow", Domestic: true},
		{Name: "Lion", Sound: "Roar", Domestic: false},
	}

	animal := Animals{Name: "Elephant", Sound: "Trumpet", Domestic: false}

	name := Animals{Name: "Rabbit"}
	newName := &Animals{Name: "Mouse", Sound: "squirt", Domestic: false}
	nameVal := *newName

	fmt.Println(animals)
	fmt.Println(animal)
	fmt.Println(name)
	fmt.Println(newName)
	fmt.Println(nameVal)
}
