// Simple queue
package queue

// Queue object
type Queue[T any] struct {
	queue []T
}

// Constructor for Queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0),
	}
}

// Returns true when queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

// Returns size of queue
func (q *Queue[T]) Size() int {
	return len(q.queue)
}

// Adds item to queue
func (q *Queue[T]) Add(t T) {
	q.queue = append(q.queue, t)
}

// Peeks at first item in queue
func (q *Queue[T]) Peek() T {
	return q.queue[0]
}

// Removes fist item and returns it
func (q *Queue[T]) Remove() T {
	ret := q.queue[0]
	q.queue = q.queue[1:]
	return ret
}
