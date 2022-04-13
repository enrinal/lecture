package main

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(val int) {
	node := &Node{Val: val}
	if l.Head == nil {
		l.Head = node
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	node.Prev = cur
}

func (l *LinkedList) Remove(val int) {
	if l.Head == nil {
		return
	}
	cur := l.Head
	for cur.Next != nil {
		if cur.Val == val {
			if cur.Prev != nil {
				cur.Prev.Next = cur.Next
			}
			if cur.Next != nil {
				cur.Next.Prev = cur.Prev
			}
			if cur == l.Head {
				l.Head = cur.Next
			}
			return
		}
		cur = cur.Next
	}
}

func (l *LinkedList) PrintForward() {
	cur := l.Head
	for cur != nil {
		print(cur.Val, " ")
		cur = cur.Next
	}
	println()
}

func (l *LinkedList) PrintBackward() {
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	for cur != nil {
		print(cur.Val, " ")
		cur = cur.Prev
	}
	println()
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	for cur != nil {
		cur.Next, cur.Prev = cur.Prev, cur.Next
		cur = cur.Prev
	}
	l.Head, cur.Next = cur, nil
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
