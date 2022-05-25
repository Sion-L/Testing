package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求元素和
func sumArr(arr *[10]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	rand.Seed(time.Now().Unix())
	a := [10]int{}
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(1000)
	}
	sum := sumArr(&a)
	fmt.Println(sum)
}
