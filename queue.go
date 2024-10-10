//go:build cname

package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, true
}

func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}

	// Enqueue elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Current Queue:", queue.items)

	// Peek at the front element
	front, ok := queue.Peek()
	if ok {
		fmt.Println("Front Element:", front)
	}

	// Dequeue elements
	dequeued, ok := queue.Dequeue()
	if ok {
		fmt.Println("Dequeued Element:", dequeued) // Output: 10
	}
	fmt.Println("Queue after Dequeue:", queue.items) // Output: [20 30]

	// Check if queue is empty
	fmt.Println("Is Queue Empty?", queue.IsEmpty()) // Output: false
}
