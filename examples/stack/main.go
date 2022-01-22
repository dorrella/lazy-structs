package main

import (
	"fmt"

	"github.com/dorrella/lazy-structs/stack"
)

func main() {
	s := stack.NewStack[int]()
	fmt.Println(s.IsEmpty())
	s.Push(3)
	s.Push(2)
	s.Push(1)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())

	fmt.Println("-----")

	s2 := stack.NewStack[string]()
	fmt.Println(s2.IsEmpty())
	s2.Push("hello")
	s2.Push("goodbye")
	s2.Push("aloha")
	fmt.Println(s2.IsEmpty())
	fmt.Println(s2.Size())
	fmt.Println(s2.Peek())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Size())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Size())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Size())
}
