package main

import "fmt"

func Unmatch(a []int) int {
	unMap := make(map[int]int)
	for _, num := range a {
		if val, exists := unMap[num]; exists {
			unMap[num] = val + 1
		} else {
			unMap[num] = 1
		}
	}
	for key, val := range unMap {
		if val != 2 {
			return key
		}
	}
	return -1
}
func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
