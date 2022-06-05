package main

import "fmt"

func main() {

	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	d := a[:2:3]
	// append 超过原有数组长度则重新分配底层数组
	d = append(d, 1, 2, 3, 4)
	fmt.Println(a, d)
	fmt.Println(&a[0], &d[0])

	// cap重新分配规律 2,4,8...
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

}
