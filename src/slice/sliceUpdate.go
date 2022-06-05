package main

import "fmt"

func main() {
	s := make([]struct {
		x int
	}, 10)

	d := s[:]

	s[1].x = 100
	d[2].x = 200
	fmt.Println(s)
	fmt.Printf("%p, %p\n", &s, &s[0])
}
