package main

import "testing"

func TestExtractMiddle(t *testing.T) {
	input := [5]string{"A", "B", "C", "D", "E"}
	got := ExtractMiddle(input)
	want := []string{"B", "C", "D"}
	if len(got) != 3 {
		t.Fatalf("got length %d, want 3", len(got))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("at index %d: got %s, want %s", i, got[i], want[i])
		}
	}
}
