package main

import "fmt"

// Queue represents a singly-linked list that holds
// values of any type.
type Queue[T any] struct {
	head *QueueMember[T]
	tail *QueueMember[T]
}

type QueueMember[T any] struct {
	val  T
	next *QueueMember[T]
}

func (q *Queue[T]) Enqueue(inputs ...T) {
	for _, input := range inputs {
		m := &QueueMember[T]{val: input} // Create new node directly

		if q.tail != nil { // If tail exists, update it
			q.tail.next = m
		}
		q.tail = m // Update tail

		if q.head == nil { // If queue was empty, update head
			q.head = m
		}
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var zeroValueOfT T
		return zeroValueOfT, false
	}

	deQdValue := q.head.val
	q.head = q.head.next

	if q.head == nil {
		// If the queue becomes empty, reset the tail as well
		q.tail = nil
	}

	return deQdValue, true
}

func (q Queue[T]) String() string {
	members := []any{}
	v := q.head
	for {
		if v == nil {
			break
		}
		members = append(members, v.val)
		v = v.next
	}

	return fmt.Sprintf("%v", members)
}

func main() {
	// q := Queue[int]{head: nil}

	// q.Enqueue(1)
	// q.Enqueue(100)
	// q.Enqueue(102)
	// q.Enqueue(3)

	q := Queue[string]{head: nil}
	q.Enqueue("Dennis", "Bundi", "Mithamo")
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
