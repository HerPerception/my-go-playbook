package main

func Number(busStops [][2]int) int {
	sum := 0
	for _, num := range busStops {
		sumslice := num[0]
		sum += sumslice
		for _, num2 := range num {
			sumslice -= num2
		}
		sum += sumslice
	}
	return sum // Good Luck!
}

// func main() {
// 	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
// 	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
// 	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
// 	fmt.Println(Number([][2]int{{0, 0}}))
// }
