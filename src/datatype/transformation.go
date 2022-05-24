package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c := float64(a*a + b*b)
	c = math.Sqrt(c)
	fmt.Println(int(c))
}
