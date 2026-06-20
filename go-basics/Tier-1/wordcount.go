package main

import (
	"strings"
)

func WordCount(sentence string) map[string]int {
	str := strings.Fields(sentence)
	count := make(map[string]int)
	for _, word := range str {
		if val, exists := count[word]; exists {
			count[word] = val + 1
		} else {
			count[word] = 1
		}
	}
	return count
}

// func main() {
// 	fmt.Println(WordCount("the cat sat on the mat the cat"))
// }
