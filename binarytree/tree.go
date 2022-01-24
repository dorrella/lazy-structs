package binarytree

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type BinaryTree[K Comparable, V any] struct {
	Root *BinaryTreeNode[K, V]
}

func NewBinaryTree[K Comparable, V any]() *BinaryTree[K, V] {
	return &BinaryTree[K, V]{nil}
}

func (bt *BinaryTree[K, V]) Add(key K, val V) {
	if bt.Root == nil {
		bt.Root = NewBinaryTreeNode(key, val)
		return
	}

	add(key, val, bt.Root)
}

func add[K Comparable, V any](key K, val V, node *BinaryTreeNode[K, V]) {
	if key < node.Key {
		if node.Left == nil {
			node.Left = NewBinaryTreeNode(key, val)
			return
		}
		add(key, val, node.Left)
	} else {
		if node.Right == nil {
			node.Right = NewBinaryTreeNode(key, val)
			return
		}
		add(key, val, node.Right)
	}
}

func (bt *BinaryTree[K, V]) Search(key K) (V, bool) {
	var ret V
	if bt.Root == nil {
		return ret, false
	}

	node := bt.Root
	for node != nil {
		if node.Key == key {
			return node.Val, true
		}

		if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return ret, false
}
