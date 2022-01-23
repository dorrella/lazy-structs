package btree

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type BTree[K Comparable, V any] struct {
	Root *BTreeNode[K, V]
}

func NewBTree[K Comparable, V any]() *BTree[K, V] {
	return &BTree[K, V]{nil}
}

func (bt *BTree[K, V]) Add(key K, val V) {
	if bt.Root == nil {
		bt.Root = NewBTreeNode(key, val)
		return
	}

	add(key, val, bt.Root)
}

func add[K Comparable, V any](key K, val V, node *BTreeNode[K, V]) {
	if key < node.Key {
		if node.Left == nil {
			node.Left = NewBTreeNode(key, val)
			return
		}
		add(key, val, node.Left)
	} else {
		if node.Right == nil {
			node.Right = NewBTreeNode(key, val)
			return
		}
		add(key, val, node.Right)
	}
}

func (bt *BTree[K, V]) Search(key K) (V, bool) {
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
