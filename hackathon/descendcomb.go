package main

import (
	"strings"
)

func DuplicateEncode(word string) string {
	myMap := make(map[rune]int)
	var build strings.Builder
	for _, ch := range word {

		if val, exists := myMap[ch]; exists {
			myMap[ch] = val + 1
		} else {
			myMap[ch] = 1
		}
	}
	for _, ch := range word {
		if myMap[ch] < 2 {
			build.WriteString("(")
		} else {
			build.WriteString(")")
		}
	}
	return build.String()
}

// func main() {
// 	fmt.Println(DuplicateEncode("(())())"))
// }
