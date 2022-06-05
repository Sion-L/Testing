package main

import "fmt"

func myTest(a [5]int, target int) {
	// 遍历数组 - 从下标0开始
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历 - 从下标1开始，当遍历的值等与8-a[i]时则打印
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
