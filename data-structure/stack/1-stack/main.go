package main

import "fmt"

type Stack struct {
	Top  int
	Size int
	Data []string
}

func NewStack(size int) *Stack {
	return &Stack{
		Top:  -1,
		Size: size,
		Data: make([]string, size),
	}
}

func (s *Stack) Push(v string) {
	if s.Top+1 == s.Size {
		fmt.Println("stack overflow")
		return
	}
	s.Top++
	s.Data[s.Top] = v
}

func (s *Stack) Pop() string {
	if s.Top == -1 {
		fmt.Println("stack underflow")
		return ""
	}
	v := s.Data[s.Top]
	s.Top--
	return v
}

func (s *Stack) Peek() string {
	if s.Top == -1 {
		fmt.Println("stack underflow")
		return ""
	}
	return s.Data[s.Top]
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

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
