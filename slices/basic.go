package main

import "fmt"

func main() {
	// No. 1: Initializing a slice with five elements - prime numbers

	slice := []int{1, 2, 3, 5, 7}
	fmt.Println(slice)

	/* No. 2: Creating an empty slice and appending elements with a loop, printing the cap
	and the len at every append*/
	var newSlice []int
	for i := 1; i <= 10; i++ {
		newSlice = append(newSlice, i)
		fmt.Println(newSlice)
		fmt.Printf("Length at iteration %d is %d\n", i, len(newSlice))
		fmt.Printf("Capacity at iteration %d is %d\n", i, cap(newSlice))
	}

	/* No. 3: Working on sub-slices*/

	Slice := []int{10, 20, 30, 40, 50}
	subSlice := Slice[1 : len(Slice)-1]
	fmt.Println(Slice)
	fmt.Println(subSlice)

}
