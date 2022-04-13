package main

import "fmt"

/*
	Stack is a data structure that follows the Last In First Out (LIFO) principle.
	Stack is a linear data structure.
	Implementation of stack using node.
	a node has a value and a pointer to the prev node
*/

type Stack struct {
	top    *Node
	length int
}

type Node struct {
	value interface{}
	prev  *Node
}

// NewStack returns a pointer to a new stack.
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Push adds a new value to the top of the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.length++
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.prev
	s.length--
	return value
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Length returns the number of items in the stack.
func (s *Stack) Length() int {
	return s.length
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// Clear removes all values from the stack.
func (s *Stack) Clear() {
	s.top = nil
	s.length = 0
}

func (s *Stack) String() string {
	str := "Stack: "
	for current := s.top; current != nil; current = current.prev {
		str += fmt.Sprintf("%v ", current.value)
	}
	return str
}

func main() {
	stacko := NewStack()
	stacko.Push("red")
	stacko.Push("blue")
	stacko.Push("green")
	stacko.Push("yellow")
	stacko.Push("red")
	stacko.Push("blue")

	fmt.Println(stacko)
	fmt.Println(stacko.Pop())
	fmt.Println(stacko.Pop())
	fmt.Println(stacko.Pop())
}
