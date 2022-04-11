package main

import "fmt"

type queue struct {
	items []interface{}
}

func NewQueue() *queue {
	return &queue{}
}

// Enqueue adds an item to the queue
func (q *queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the queue
func (q *queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *queue) Size() int {
	return len(q.items)
}

func (q *queue) Clear() {
	q.items = []interface{}{}
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
