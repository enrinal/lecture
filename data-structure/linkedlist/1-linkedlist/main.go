package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (l *LinkedList) Add(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
	} else {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (l *LinkedList) Remove(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}
	cur := l.Head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}
}

func (l *LinkedList) PrintForward() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func (l *LinkedList) PrintBackward() {
	l.Reverse()
	l.PrintForward()
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}
	cur := l.Head
	var prev *Node
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}

func main() {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.PrintForward()
	l.PrintBackward()
	l.Remove(3)
	l.PrintForward()
	l.PrintBackward()
	l.Reverse()
	l.PrintForward()
	l.PrintBackward()
}
