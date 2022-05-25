package main

import "fmt"

func testarr(s [2]int) {
	fmt.Printf("s: %p\n", &s)
	s[1] = 1000
}

func main() {

	arr1 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr1)

	arr2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)

	// a的值传给了s，但是两者指针不一样，值一样，是值拷贝
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)

	testarr(a)
	fmt.Println(a)
}
