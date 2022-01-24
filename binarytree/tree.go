// Basic binary tree implrementation
package binarytree

// Interface for comparable number constraing
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// BinaryTree Struct
// Simple wrapper around Root Node
type BinaryTree[K Comparable, V any] struct {
	Root *BinaryTreeNode[K, V]
}

// Basic constructor for BinaryTree
func NewBinaryTree[K Comparable, V any]() *BinaryTree[K, V] {
	return &BinaryTree[K, V]{nil}
}

// Add key/val pair to BinaryTree
//
// Calls wrapper around NewBinaryTreeNode and
// adds it to the BinaryTree
func (bt *BinaryTree[K, V]) Add(key K, val V) {
	if bt.Root == nil {
		bt.Root = NewBinaryTreeNode(key, val)
		return
	}

	add(key, val, bt.Root)
}

// Recursive helper function for add
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

// Search for key in BinaryTree
//
// Returns `value, true` if found `_, false` if not
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
