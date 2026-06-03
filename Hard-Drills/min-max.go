package main

func FindMinMax(myMap map[string]int) (string, string) {
	max := 0
	min := max
	minVal := ""
	maxVal := ""
	for key, val := range myMap {
		if min == 0 {
			min = val
			minVal = key
		}
		if min > val {
			min = val
			minVal = key
		}
		if max < val {
			max = val
			maxVal = key
		}
	}
	return maxVal, minVal
}

// func main() {
// 	inventoryMap := map[string]int{"pens": 15, "books": 120, "erasers": 5, "rulers": 50, "crayons": 0}
// 	min, max := FindMinMax(inventoryMap)
// 	fmt.Printf("%q, %q", min, max)
// }
