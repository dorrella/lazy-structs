package stack

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		stack: make([]T, 0),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}

func (s *Stack[T]) Push(t T) {
	s.stack = append(s.stack, t)
}

func (s *Stack[T]) Peek() T {
	l := len(s.stack) - 1
	return s.stack[l]
}

func (s *Stack[T]) Pop() T {
	l := len(s.stack) - 1
	ret := s.stack[l]
	s.stack = s.stack[:l]
	return ret
}
