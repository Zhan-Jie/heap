package main

import (
	"fmt"

	heap "./src/heap"
)

// MyInt wrapper
type MyInt int

// GreaterThan implements interface
func (i MyInt) GreaterThan(b *heap.Comparable) bool {
	bInt, ok := (*b).(MyInt)
	if ok {
		return i > MyInt(bInt)
	}
	panic(fmt.Sprintf("%v is not a int\n", *b))
}

func main() {
	h := heap.Heap{}
	h.Add(MyInt(1))
	h.Add(MyInt(1))
	h.Add(MyInt(2))
	h.Add(MyInt(2))
	h.Add(MyInt(5))
	h.Add(MyInt(5))
	h.Add(MyInt(4))
	h.Add(MyInt(4))
	fmt.Printf("hello, sb~, %v\n", h)
	
	for {
		ele := h.Remove()
		if (ele == nil) {
			return
		}
		fmt.Printf("pop: %v\n", ele)
	}
}
