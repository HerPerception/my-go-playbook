package main

func CelsiusToFahrenheit(c float64) float64 {
	var fareinheit float64
	fareinheit = (c * 9 / 5) + 32
	return fareinheit
}

// func main() {
// 	c := float64(100)
// 	fareinheit := CelsiusToFahrenheit(c)
// 	fmt.Printf("%.2f = %.2f", c, fareinheit)
// }
