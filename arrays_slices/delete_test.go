package main

import "testing"

func TestRemoveAtIndex(t *testing.T) {
	colors := []string{"red", "green", "blue", "yellow"}
	// Delete "blue" at index 2
	got := RemoveAtIndex(colors, 2)
	want := []string{"red", "green", "yellow"}
	if len(got) != 3 {
		t.Fatalf("expected length 3, got %d", len(got))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("at index %d: got %s, want %s", i, got[i], want[i])
		}
	}
}
