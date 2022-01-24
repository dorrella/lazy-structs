package binarytree

// Node for Binary Tree
type BinaryTreeNode[K Comparable, V any] struct {
	Key   K
	Val   V
	Left  *BinaryTreeNode[K, V]
	Right *BinaryTreeNode[K, V]
}

// Simple Constructur
//
// Should probably use BinaryTree.Add
func NewBinaryTreeNode[K Comparable, V any](key K, val V) *BinaryTreeNode[K, V] {
	return &BinaryTreeNode[K, V]{
		Key:   key,
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
