package main

import "fmt"

func main() {

	// append 可为slice添加slice和元素
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := append(a, b...)
	d := append([]byte("hello "), "world"...)
	fmt.Printf("%#v\n,%s", c, d)

}
