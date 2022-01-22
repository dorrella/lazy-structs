package main

import (
	"fmt"

	"github.com/dorrella/lazy-structs/linkedlist"
)

func PrintList[T any](l *linkedlist.LinkedList[T]) {
	n := l.Head
	for n != nil {
		fmt.Printf("%v ", n.Val)
		n = n.Next
	}
	fmt.Println()

}

func PrintReversed[T any](l *linkedlist.DLinkedList[T]) {
	n := l.Tail
	for n != nil {
		fmt.Printf("%v ", n.Val)
		n = n.Last
	}
	fmt.Println()
}

func main() {
	in := []string{"hello", "world", "aloha"}

	ll := linkedlist.NewLinkedList[string]()
	ll.InitList(in)
	PrintList(ll)

	dl := linkedlist.NewDLinkedList[string]()
	dl.InitList(in)
	PrintReversed(dl)

}
