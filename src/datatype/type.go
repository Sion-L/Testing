package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		071根据科学计数法，有效数字为71，且为8进制(071中没有大于8的数)，转化成10进制为57
		0x1F则为16进制(1,2,3,4,5,6,7,8,9,10,a,b,c,d,e,f)，转换成10进制为31
		1e9为1*10的九次方，科学计数法为1e+09
		math.MaxInt8可计算int8的最大值
	*/
	a, b, c, d := 071, 0x1F, 1e9, math.MaxInt8
	var x complex128 = complex(1, 2)
	fmt.Println(a, b, c, d, x)

	fmt.Println("\"test\"") // "test"

	s := "abcd"
	t := "efg"
	fmt.Println(s + t) // abcdefg

}
