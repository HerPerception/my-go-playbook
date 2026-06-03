package main

func RemoveDuplicates[T comparable](items []T) []T {
	uniqueValues := make(map[T]int)
	var uniqueSlice []T
	for i, key := range items {
		if _, exists := uniqueValues[key]; exists {
			continue
		} else {
			uniqueSlice = append(uniqueSlice, key)
			uniqueValues[key] = i
		}
	}
	return uniqueSlice
}

// func main() {
// 	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 1}))
// 	fmt.Println(RemoveDuplicates([]string{"go", "js", "go", "js"}))
// 	fmt.Println(RemoveDuplicates([]bool{false, true, false}))
// 	fmt.Println(RemoveDuplicates([]string{"apple", "banana", "apple", "cherry", "banana"}))

// }
