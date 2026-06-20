package main

import "fmt"

func InitializeArray() [4]int {
	var arr [4]int
	for i := range arr {
		arr[i] = i * 10
	}
	return arr
}

func main() {
	fmt.Println(ExtractMiddle([5]string{"a", "b", "c", "d", "e"}))
}
