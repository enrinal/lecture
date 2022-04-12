package main

import "fmt"

/*
	Stack is a data structure that follows the Last In First Out (LIFO) principle.
	Stack is a linear data structure.
	Implementation of stack using array.
	Stack overflow: when the stack is full.
	Stack underflow: when the stack is empty.
*/

type Stack struct {
	Top  int
	Size int
	Data []string
}

// NewStack returns a new stack.
func NewStack(size int) *Stack {
	return &Stack{
		Top:  -1,
		Size: size,
		Data: make([]string, size),
	}
}

// Push adds a new value to the top of the stack.
func (s *Stack) Push(v string) {
	if s.Top+1 == s.Size {
		fmt.Println("stack overflow")
		return
	}
	s.Top++
	s.Data[s.Top] = v
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack) Pop() string {
	if s.Top == -1 {
		fmt.Println("stack underflow")
		return ""
	}
	v := s.Data[s.Top]
	s.Top--
	return v
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack) Peek() string {
	if s.Top == -1 {
		fmt.Println("stack underflow")
		return ""
	}
	return s.Data[s.Top]
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

// Clear removes all values from the stack.
func (s *Stack) Clear() {
	s.Top = -1
}

func (s *Stack) String() string {
	str := "Stack: "
	for i := 0; i <= s.Top; i++ {
		str += s.Data[i] + " "
	}
	return str
}

func main() {
	stacko := NewStack(5)
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
	fmt.Println(stacko.Pop())
	fmt.Println(stacko.Pop())
	fmt.Println(stacko.Pop())
}
