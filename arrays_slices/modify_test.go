package main

import "testing"

func TestModifySliceAndArray(t *testing.T) {
	originalArray := [3]int{1, 2, 3}
	// Create a slice pointing directly to the original array
	targetSlice := originalArray[:]
	ModifySliceAndArray(targetSlice, 999)
	// Verify the underlying array mutated without referencing the slice directly
	if originalArray[0] != 999 {
		t.Errorf("expected originalArray[0] to change to 999, got %d", originalArray[0])
	}
}
