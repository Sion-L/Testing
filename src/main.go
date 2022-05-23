package main

import (
	_ "Once/src/hello"
	"fmt"
	"os"
)

func main() {
	// 1.如果要使用hello包中的Print函数，则把_去掉
	// 2.下划线在代码中的使用 - 忽略返回的参数
	buf := make([]byte, 1024)
	f, _ := os.Open("./test.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			fmt.Println("file is not exsit")
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
