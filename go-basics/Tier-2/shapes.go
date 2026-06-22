package main

import "fmt"

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

const Pi = 3.14159265358979323846

func (c Circle) Area() float64 {
	rad := c.radius * c.radius
	return Pi * rad
}

func (c Circle) Perimeter() float64 {
	per := 2 * Pi * c.radius
	return per
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func PrintShape(s Shape) {
	fmt.Printf("The Area of this shape is %.2fcm² and the Perimeter is %.2fcm\n", s.Area(), s.Perimeter())
}
func TaskOne() {
	c := Circle{radius: 5}
	r := Rectangle{width: 10, height: 5}

	PrintShape(c)
	PrintShape(r)
}
