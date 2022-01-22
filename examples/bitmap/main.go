package main

import (
	"fmt"

	"github.com/dorrella/lazy-structs/bitmap"
)

func main() {
	bm := bitmap.NewBitMap()

	fmt.Println(bm.Get(0))
	fmt.Println(bm.Get(10000))

	bm.Set(0, true)
	bm.Set(10000, true)

	fmt.Println(bm.Get(0))
	fmt.Println(bm.Get(10000))
}
