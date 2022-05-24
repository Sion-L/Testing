package main

import "fmt"

const (
	n1 = 100
	n2
	n3
)

/*
iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计
数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用
*/
const (
	t  = iota // 0
	t1        // 1
	t2        // 2
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const n5 = iota //0

/*
1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也
就是由10变成了1000，也就是十进制的8。
*/
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(t, t1, t2)
	fmt.Println(KB, MB, GB, TB, PB)
}
