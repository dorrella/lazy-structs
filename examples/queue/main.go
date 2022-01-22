package main

import (
	"fmt"

	"github.com/dorrella/lazy-structs/queue"
)

func main() {
	q1 := queue.NewQueue[int]()

	fmt.Println(q1.IsEmpty())
	q1.Add(3)
	q1.Add(2)
	q1.Add(1)
	fmt.Println(q1.IsEmpty())
	fmt.Println(q1.Size())
	fmt.Println(q1.Peek())
	fmt.Println(q1.Remove())
	fmt.Println(q1.Size())
	fmt.Println(q1.Remove())
	fmt.Println(q1.Size())
	fmt.Println(q1.Remove())
	fmt.Println(q1.Size())

	fmt.Println("-----")

	q2 := queue.NewQueue[string]()
	fmt.Println(q2.IsEmpty())
	q2.Add("hello")
	q2.Add("goodbye")
	q2.Add("aloha")
	fmt.Println(q2.IsEmpty())
	fmt.Println(q2.Size())
	fmt.Println(q2.Peek())
	fmt.Println(q2.Remove())
	fmt.Println(q2.Size())
	fmt.Println(q2.Remove())
	fmt.Println(q2.Size())
	fmt.Println(q2.Remove())
	fmt.Println(q2.Size())
}
