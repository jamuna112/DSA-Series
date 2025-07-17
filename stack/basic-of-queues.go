package stack

import "fmt"

type Queue struct {
	items []int
}

// Push operation
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Pop operation
func (q *Queue) Dequeue() int {

	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item

}

// Peek operation - returning the first value in the queue
func (q *Queue) Front() int {

	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) PrintQueue() {
	for _, que := range q.items {
		fmt.Print(que, " ")
	}
}
