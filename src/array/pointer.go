package main

import "fmt"

func pointArr(arr *[5]int) {
	arr[0] = 100
	fmt.Println(arr)
}

func main() {
	// 指针数组：数组成员类型都是指针[n]*T
	// 数组指针：指向数组的指针*[n]T

	arr1 := [5]int{1, 2, 3, 4, 5}
	pointArr(&arr1) // 传递数组指针实现值传递
	fmt.Println(arr1)

}
