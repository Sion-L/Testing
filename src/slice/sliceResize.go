package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n, cap(a) : %v\n", a, len(a), cap(a))
	b := a[1:2] // cap = 原切片底层数组长度 - 起始位1 = 4 - 1 = 3
	fmt.Printf("slice b : %v , len(b) : %v\n, cap(b) : %v\n", b, len(b), cap(b))
	c := b[0:3] // cap = 切片b底层数组长度 - 起始位0 = 3 - 0 = 3
	fmt.Printf("slice c : %v , len(c) : %v\n, cap(c): %v\n", c, len(c), cap(c))
}
