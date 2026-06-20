package main

func RotateLeft(a [5]int) [5]int {
	var b [5]int
	for i := range a {
		if i == len(a)-1 {
			b[i] = a[0]
		} else {
			b[i] = a[i+1]
		}
	}
	return b
}

// func main() {
// 	fmt.Println(RotateLeft([5]int{1, 2, 3, 4, 5}))
// }
