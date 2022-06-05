package main

import "fmt"

func main() {

	s := []int{1, 2, 3}
	t := s[1:3]
	fmt.Println(t)

	// 初始化
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := arr1[1:3]
	m := arr1[:6]
	e := arr1[1:]
	i := arr1[:]
	o := arr1[1:3:6]
	fmt.Println(n, m, e, i, o)

}
