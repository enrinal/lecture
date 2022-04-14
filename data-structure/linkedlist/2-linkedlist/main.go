package main

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *LinkedList) PushBack(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *LinkedList) PushFront(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}
	node.Next = l.Head
	l.Head = node
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

func (l *LinkedList) Print() {
	cur := l.Head
	for cur != nil {
		print(cur.Val, " ")
		cur = cur.Next
	}
	println()
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}

	var prev *Node
	cur := l.Head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func main() {
	l := NewLinkedList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.Print()

	l.PushFront(0)

	l.Print()
	//l.PushFront(0)
	//
	//l.Print()
	//l.Remove(3)
	//l.Print()
	//l.Reverse()
	//l.Print()
}
