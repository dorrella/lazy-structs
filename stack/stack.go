// Simple stack
package stack

// Stack struct
type Stack[T any] struct {
	stack []T
}

// Constructor for stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		stack: make([]T, 0),
	}
}

// Returns true if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

// Returns size of stack
func (s *Stack[T]) Size() int {
	return len(s.stack)
}

// Pushes item onto stack
func (s *Stack[T]) Push(t T) {
	s.stack = append(s.stack, t)
}

// Peeks at top item and returns it
func (s *Stack[T]) Peek() T {
	l := len(s.stack) - 1
	return s.stack[l]
}

// Removes first item from stack and returns it
func (s *Stack[T]) Pop() T {
	l := len(s.stack) - 1
	ret := s.stack[l]
	s.stack = s.stack[:l]
	return ret
}
