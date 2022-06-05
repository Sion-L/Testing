package main

import "fmt"

func main() {
	// slice & struct 可直接修改struct slice 成员
	s := make([]struct {
		x string
	}, 10)

	d := s[:]
	d[1].x = "li"
	d[2].x = "lang"
	fmt.Printf("%s\n,%s", s, d)

}
