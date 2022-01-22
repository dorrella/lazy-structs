package queue

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0),
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}

func (q *Queue[T]) Add(t T) {
	q.queue = append(q.queue, t)
}

func (q *Queue[T]) Peek() T {
	return q.queue[0]
}

func (q *Queue[T]) Remove() T {
	ret := q.queue[0]
	q.queue = q.queue[1:]
	return ret
}
