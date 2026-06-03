package main

func FilterEvens(nums []int) []int {
	var evenSlice []int
	for _, num := range nums {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}
	return evenSlice
}

// func main() {
// 	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
// 	fmt.Println(FilterEvens(num))
// }
