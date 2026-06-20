package main

import "testing"

func TestBuildAndGrow(t *testing.T) {
	length := 2
	capacity := 4
	extra := 99
	got := BuildAndGrow(length, capacity, extra)
	// Expected length is initial length (2) + appended item (1) = 3
	if len(got) != 3 {
		t.Errorf("expected length 3, got %d", len(got))
	}
	if got[2] != 99 {
		t.Errorf("expected last element to be 99, got %d", got[2])
	}
}
