package main

import "fmt"

func main() {

	arr1 := [2]int{}
	fmt.Println(len(arr1), cap(arr1))

	arr2 := [...][2]int{{1, 2}, {3, 4}}
	for _, v1 := range arr2 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
