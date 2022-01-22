package linkedlist

type LinkNode[T any] struct {
	Val  T
	Next *LinkNode[T]
}

func NewLinkNode[T any](val T) *LinkNode[T] {
	return &LinkNode[T]{
		Val:  val,
		Next: nil,
	}
}

type DLinkNode[T any] struct {
	Val  T
	Next *DLinkNode[T]
	Last *DLinkNode[T]
}

func NewDLinkNode[T any](val T) *DLinkNode[T] {
	return &DLinkNode[T]{
		Val:  val,
		Next: nil,
		Last: nil,
	}
}
