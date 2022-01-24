package linkedlist

// Doubly linked list.
type DLinkedList[T any] struct {
	Head *DLinkNode[T]
	Tail *DLinkNode[T]
}

// Simple constructer for DLinkedList
func NewDLinkedList[T any]() *DLinkedList[T] {
	return &DLinkedList[T]{
		Head: nil,
		Tail: nil,
	}
}

// Initializes linked list to array.
// overwrites Head and Tail
func (ll *DLinkedList[T]) InitList(in []T) {
	ll.Head = nil
	ll.Tail = nil
	if len(in) == 0 {
		return
	}

	ll.Head = NewDLinkNode[T](in[0])
	node := ll.Head

	for i := 1; i < len(in); i++ {
		node.Next = NewDLinkNode[T](in[i])
		node.Next.Last = node
		node = node.Next
	}

	ll.Tail = node

}
