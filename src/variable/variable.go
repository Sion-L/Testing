package main

import "fmt"

func foo() (int, string) {
	return 1, "foo"
}

func main() {

	test, test1 := "test", "test1"
	fmt.Println(test, test1)

	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)

}
