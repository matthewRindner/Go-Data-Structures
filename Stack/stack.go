package main

import "fmt"

//stack represent stack that holds a slice
type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) pop() int {
	lengthOfStack := len(s.items) - 1
	toRemove := s.items[lengthOfStack]
	s.items = s.items[:lengthOfStack]
	return toRemove
}

func main() {
	myStack := Stack{}

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	fmt.Println(myStack)

	myStack.pop()
	fmt.Println(myStack)
}
