package main

import "fmt"

func main() {
	// make
	s := make([]int, 10)
	m := make([]int, 10, 12)
	fmt.Printf("%#v\n", s, m)

	a := []int{1, 2, 3, 4}
	b := a[1:3]
	b[0] += 100
	b[1] += 100
	fmt.Println(a, b)

	// 切片不指定cap，则cap=len
	w := []int{1, 2, 3, 8: 10}
	e := make([]int, 10)
	fmt.Printf("%#v\n", w, e, len(e), cap(e))

}
