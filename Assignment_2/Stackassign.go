package main

import "fmt"

//Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end (top)
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end (top)
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(151)
	myStack.Push(251)
	myStack.Push(351)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}