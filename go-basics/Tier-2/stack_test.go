package main

import (
	"testing"
)

func TestStack_PushAndIsEmpty(t *testing.T) {
	t.Run("New stack is empty", func(t *testing.T) {
		s := &Stack{}
		if !s.IsEmpty() {
			t.Error("Expected new stack to be empty, but IsEmpty() returned false")
		}
	})

	t.Run("Push items clears IsEmpty", func(t *testing.T) {
		s := &Stack{}
		s.Push(10)
		if s.IsEmpty() {
			t.Error("Expected stack to not be empty after pushing an item")
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("Peek empty stack returns error", func(t *testing.T) {
		s := &Stack{}
		_, err := s.Peek()
		if err == nil {
			t.Error("Expected error when peering into an empty stack, got nil")
		}
	})

	t.Run("Peek returns top item without removing it", func(t *testing.T) {
		s := &Stack{}
		s.Push(100)
		s.Push(200)

		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Unexpected error on Peek: %v", err)
		}
		if val != 200 {
			t.Errorf("Expected Peek to return 200, got %d", val)
		}

		// Double check that size did not change
		val2, _ := s.Peek()
		if val2 != 200 {
			t.Errorf("Expected second Peek to still return 200, got %d", val2)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("Pop empty stack returns error", func(t *testing.T) {
		s := &Stack{}
		_, err := s.Pop()
		if err == nil {
			t.Error("Expected error when popping from an empty stack, got nil")
		}
	})

	t.Run("Pop removes items in Last-In-First-Out order", func(t *testing.T) {
		s := &Stack{}
		s.Push(5)
		s.Push(10)

		val1, err1 := s.Pop()
		if err1 != nil || val1 != 10 {
			t.Errorf("First Pop expected 10, got %d (error: %v)", val1, err1)
		}

		val2, err2 := s.Pop()
		if err2 != nil || val2 != 5 {
			t.Errorf("Second Pop expected 5, got %d (error: %v)", val2, err2)
		}

		if !s.IsEmpty() {
			t.Error("Expected stack to be empty after popping all elements")
		}
	})
}

func TestStack_MainWorkflowSequence(t *testing.T) {
	t.Run("Simulate task step-by-step main lifecycle sequence", func(t *testing.T) {
		s := &Stack{}

		// 1. Pushes 1, 2, 3
		s.Push(1)
		s.Push(2)
		s.Push(3)

		// 2. Pops twice
		p1, err := s.Pop()
		if err != nil || p1 != 3 {
			t.Errorf("First sequence Pop expected 3, got %d, error: %v", p1, err)
		}

		p2, err := s.Pop()
		if err != nil || p2 != 2 {
			t.Errorf("Second sequence Pop expected 2, got %d, error: %v", p2, err)
		}

		// 3. Peeks
		pk, err := s.Peek()
		if err != nil || pk != 1 {
			t.Errorf("Sequence Peek expected 1, got %d, error: %v", pk, err)
		}

		// 4. Checks IsEmpty()
		if s.IsEmpty() {
			t.Error("Expected stack to contain remaining item '1', but IsEmpty returned true")
		}

		// Clear out final item to ensure closure consistency
		p3, err := s.Pop()
		if err != nil || p3 != 1 {
			t.Errorf("Final cleanup Pop expected 1, got %d, error: %v", p3, err)
		}

		if !s.IsEmpty() {
			t.Error("Expected stack to be fully empty at final tracking check")
		}
	})
}
