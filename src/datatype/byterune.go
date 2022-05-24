package main

import "fmt"

// 遍历字符串
func travesalString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

func main() {

	/*
			go语言字符有两种：
			uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
		    rune类型，代表一个 UTF-8字符。
	*/

	travesalString("ni hao ma")
}
