package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("Empty, nothing to pop!")
	}
	temp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return temp, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("Empty, nothing to peek at!")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
func TaskFour() {
	stack := Stack{items: []int{7, 8}}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())

}
