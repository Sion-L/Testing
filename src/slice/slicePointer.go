package main

import "fmt"

func main() {
	// 用指针修改底层数组的值来修改原切片的值
	s := []int{1, 2, 3, 4}
	p := &s[2]
	*p = 8
	fmt.Printf("%#v\n", s, *p)

	// [][]int，是指元素类型为[]int
	t := [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3},
		[]int{4, 5},
	}
	fmt.Println(t)
}
