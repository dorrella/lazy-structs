// Basic single and doubly linked lists
package linkedlist

// Simple node for LinkedList
type LinkNode[T any] struct {
	Val  T
	Next *LinkNode[T]
}

// Constructer for LinkNode
func NewLinkNode[T any](val T) *LinkNode[T] {
	return &LinkNode[T]{
		Val:  val,
		Next: nil,
	}
}

// Simple node for Doubly LinkedList
type DLinkNode[T any] struct {
	Val  T
	Next *DLinkNode[T]
	Last *DLinkNode[T]
}

// Constructor for DLinkNode
func NewDLinkNode[T any](val T) *DLinkNode[T] {
	return &DLinkNode[T]{
		Val:  val,
		Next: nil,
		Last: nil,
	}
}
