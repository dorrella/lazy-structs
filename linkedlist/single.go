package linkedlist

type LinkedList[T any] struct {
	Head *LinkNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
	}
}

func (ll *LinkedList[T]) InitList(in []T) {
	ll.Head = nil
	if len(in) == 0 {
		return
	}

	ll.Head = NewLinkNode[T](in[0])
	node := ll.Head

	for i := 1; i < len(in); i++ {
		node.Next = NewLinkNode[T](in[i])
		node = node.Next
	}

}
