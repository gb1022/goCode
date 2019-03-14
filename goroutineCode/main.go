package main

import (
	"fmt"
	"time"
)

var a [10]int

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//				fmt.Printf("a[%v]:%v\n", i, a[i])
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

	我的测试 := a
	fmt.Println("...", 我的测试)
	加数1 := 1
	加数2 := 2
	fmt.Println("和等于：", 加法(加数1, 加数2))
}

func 加法(a int, b int) int {
	return a + b
}
