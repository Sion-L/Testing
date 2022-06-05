package main

import "fmt"

func main() {
	s := [5]struct {
		x int
	}{}

	d := s[:]

	s[1].x = 100
	d[2].x = 200
	fmt.Println(s)
	fmt.Printf("%p, %p\n", &s, &s[0])
}
