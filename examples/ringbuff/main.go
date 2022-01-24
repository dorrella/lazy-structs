package main

import (
	"fmt"

	rb "github.com/dorrella/lazy-structs/ringbuffer"
)

func PrintRB[T any](buff *rb.RingBuffer[T]) {
	l := buff.Size()
	if l == 0 {
		return
	}

	for i := 0; i < l; i++ {
		obj := buff.Get(i)
		fmt.Printf("%v ", obj)
	}
	fmt.Println()
}

func main() {
	buff := rb.NewRingBuffer[int](3)
	PrintRB[int](buff)
	fmt.Println("----")
	buff.Add(1)
	buff.Add(2)
	PrintRB[int](buff)
	fmt.Println("----")
	buff.Add(3)
	buff.Add(4)
	PrintRB[int](buff)
}
