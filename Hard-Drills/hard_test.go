package main

import (
	"reflect"
	"sort"
	"testing"
)

// ============================================================================
// DRILL 1: The Slice Filter
// ============================================================================
func TestDrill7_FilterEvens(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Standard 1 to 20",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
		{
			name:     "All odds",
			input:    []int{1, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "Empty input slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Negative and zero values",
			input:    []int{-4, -3, 0, 3, 4},
			expected: []int{-4, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterEvens(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s Failed: Got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// DRILL 2: The Registry Check
// ============================================================================
func TestDrill8_CheckUser(t *testing.T) {
	registry := map[string]bool{"alice": true, "bob": false}

	tests := []struct {
		name       string
		user       string
		wantExists bool
		wantActive bool
	}{
		{"Existing active user", "alice", true, true},
		{"Existing inactive user", "bob", true, false},
		{"Missing user", "eve", false, false},
		{"Empty string lookup", "", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists, active := CheckUser(registry, tt.user)
			if exists != tt.wantExists || active != tt.wantActive {
				t.Errorf("%s Failed: Got (exists=%t, active=%t), want (exists=%t, active=%t)",
					tt.name, exists, active, tt.wantExists, tt.wantActive)
			}
		})
	}
}

// ============================================================================
// DRILL 3: Unique Item Finder
// ============================================================================
func TestDrill9_RemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Standard duplicates",
			input:    []string{"apple", "banana", "apple", "cherry", "banana"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "All duplicates identical",
			input:    []string{"apple", "apple", "apple"},
			expected: []string{"apple"},
		},
		{
			name:     "Empty string slice",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveDuplicates(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			// Map tracking makes slice ordering non-deterministic.
			// Sort both arrays before running your assertion.
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s Failed: Got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// DRILL 4: Grouping Data
// ============================================================================
func TestDrill10_GroupEmployees(t *testing.T) {
	tests := []struct {
		name     string
		input    []Employee
		expected map[string][]string
	}{
		{
			name: "Standard structure setup",
			input: []Employee{
				{Name: "Alice", Department: "HR"},
				{Name: "Bob", Department: "IT"},
				{Name: "Charlie", Department: "IT"},
				{Name: "David", Department: "HR"},
			},
			expected: map[string][]string{
				"HR": {"Alice", "David"},
				"IT": {"Bob", "Charlie"},
			},
		},
		{
			name:     "Empty staff roster",
			input:    []Employee{},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupEmployees(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			// Sort slice internals to guarantee order stability
			for k := range result {
				sort.Strings(result[k])
			}
			for k := range tt.expected {
				sort.Strings(tt.expected[k])
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s Failed: Got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// DRILL 5: The LRU Cache Tracker
// ============================================================================
func TestDrill11_LRUCache(t *testing.T) {
	t.Run("Standard history eviction", func(t *testing.T) {
		cache := &LRUCache{}
		cache.Visit("home")
		cache.Visit("about")
		cache.Visit("contact")
		cache.Visit("pricing") // Overflows limit max capacity (3)

		expected := []string{"about", "contact", "pricing"}
		if !reflect.DeepEqual(cache.History, expected) {
			t.Errorf("Drill 5 Eviction Failed: Got %v, want %v", cache.History, expected)
		}
	})

	t.Run("Under capacity limits", func(t *testing.T) {
		cache := &LRUCache{}
		cache.Visit("home")
		cache.Visit("about")

		expected := []string{"home", "about"}
		if !reflect.DeepEqual(cache.History, expected) {
			t.Errorf("Drill 5 Under-capacity Failed: Got %v, want %v", cache.History, expected)
		}
	})
}

// ============================================================================
// DRILL 6: Inventory Min/Max Finder
// ============================================================================
func TestDrill12_FindMinMax(t *testing.T) {
	tests := []struct {
		name      string
		inventory map[string]int
		wantMax   string
		wantMin   string
	}{
		{
			name:      "Standard balance lists",
			inventory: map[string]int{"pens": 15, "books": 120, "erasers": 5, "rulers": 50},
			wantMax:   "books",
			wantMin:   "erasers",
		},
		{
			name:      "Single entity registration",
			inventory: map[string]int{"solo-item": 42},
			wantMax:   "solo-item",
			wantMin:   "solo-item",
		},
		{
			name:      "Empty Map initialization",
			inventory: map[string]int{},
			wantMax:   "",
			wantMin:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxProd, minProd := FindMinMax(tt.inventory)
			if maxProd != tt.wantMax || minProd != tt.wantMin {
				t.Errorf("%s Failed: Got Max=%q Min=%q, want Max=%q Min=%q",
					tt.name, maxProd, minProd, tt.wantMax, tt.wantMin)
			}
		})
	}
}
