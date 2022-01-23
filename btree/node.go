package btree

type BTreeNode[K Comparable, V any] struct {
	Key   K
	Val   V
	Left  *BTreeNode[K, V]
	Right *BTreeNode[K, V]
}

func NewBTreeNode[K Comparable, V any](key K, val V) *BTreeNode[K, V] {
	return &BTreeNode[K, V]{
		Key:   key,
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
