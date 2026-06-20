package main

func BuildAndGrow(length int, capacity int, extra int) []int {
	slice := make([]int, length, capacity)
	return append(slice, extra)
}
