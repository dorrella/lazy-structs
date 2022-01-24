package main

import (
	"fmt"

	"github.com/dorrella/lazy-structs/binarytree"
)

func main() {
	bt := binarytree.NewBinaryTree[int, string]()
	bt.Add(5, "V")
	bt.Add(9, "IX")
	bt.Add(1, "I")
	bt.Add(3, "III")
	bt.Add(8, "VIII")
	bt.Add(4, "IV")

	v, ok := bt.Search(4)
	if ok {
		fmt.Println(v)
	}
}
