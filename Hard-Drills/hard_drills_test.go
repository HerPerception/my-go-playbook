package main

import (
	"reflect"
	"testing"
)

func TestDrill1_FilterEvens(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	result := FilterEvens(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drill 1 Failed: Got %v, want %v", result, expected)
	}
}

func TestDrill2_CheckUser(t *testing.T) {
	registry := map[string]bool{"alice": true, "bob": false}

	// Test existing active user
	exists, active := CheckUser(registry, "alice")
	if !exists || !active {
		t.Errorf("Drill 2 Active User Failed")
	}

	// Test missing user
	exists, _ = CheckUser(registry, "eve")
	if exists {
		t.Errorf("Drill 2 Missing User Failed")
	}
}

func TestDrill3_RemoveDuplicates(t *testing.T) {
	input := []string{"apple", "banana", "apple", "cherry", "banana"}
	expected := []string{"apple", "banana", "cherry"}
	result := RemoveDuplicates(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drill 3 Failed: Got %v, want %v", result, expected)
	}
}

func TestDrill4_GroupEmployees(t *testing.T) {
	input := []Employee{
		{Name: "Alice", Department: "HR"},
		{Name: "Bob", Department: "IT"},
		{Name: "Charlie", Department: "IT"},
		{Name: "David", Department: "HR"},
	}
	expected := map[string][]string{
		"HR": {"Alice", "David"},
		"IT": {"Bob", "Charlie"},
	}
	result := GroupEmployees(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drill 4 Failed: Got %v, want %v", result, expected)
	}
}

func TestDrill5_LRUCache(t *testing.T) {
	cache := &LRUCache{}
	cache.Visit("home")
	cache.Visit("about")
	cache.Visit("contact")
	cache.Visit("pricing") // Should push out "home"

	expected := []string{"about", "contact", "pricing"}
	if !reflect.DeepEqual(cache.History, expected) {
		t.Errorf("Drill 5 Failed: Got %v, want %v", cache.History, expected)
	}
}

// func TestDrill6_FindMinMax(t *testing.T) {
// 	inventory := map[string]int{"pens": 15, "books": 120, "erasers": 5, "rulers": 50}
// 	maxProd, minProd := FindMinMax(inventory)

// 	if maxProd != "books" || minProd != "erasers" {
// 		t.Errorf("Drill 6 Failed: Got Max=%s, Min=%s", maxProd, minProd)
// 	}
// }
