package main

func Unique(items []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, each := range items {
		if found := seen[each]; !found {
			result = append(result, each)
			seen[each] = true
		}
	}
	return result
}

// func main() {
// 	fmt.Println(Unique([]string{"go", "rust", "go", "python", "rust"}))
// }
