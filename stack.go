//go:build cname

package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	fmt.Println(stack)

	top, ok := stack.Peek()
	if ok {
		fmt.Println(top)
	}

	pop, ok := stack.Pop()
	if ok {
		fmt.Println(pop)
	}

	fmt.Println("Is stack empty?", stack.Empty())
}
