package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("can not divide by zero")
	}
	return a / b, nil
}

// func main() {
// 	div, err := Divide(0, 6)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(div)
// }
