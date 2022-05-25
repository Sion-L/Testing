package main

import "fmt"

func main() {

	var arr1 = [4]string{"ni", "hao", "cai", "a"}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]struct {
		name string
		age  int
	}{
		{"lilang", 18},
		{"lisi", 60},
	}

	fmt.Println(arr1, arr2, arr3)
}
