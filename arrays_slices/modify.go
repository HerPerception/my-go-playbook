package main

func ModifySliceAndArray(slice []int, newValue int) {
	slice[0] = newValue
	//slice = append(slice, newValue)
}

// func main() {
// 	originalArray := [3]int{1, 2, 3}
// 	targetSlice := originalArray[:]
// 	ModifySliceAndArray(targetSlice, 999)
// 	fmt.Println(originalArray)
// 	fmt.Println(targetSlice)
// }
