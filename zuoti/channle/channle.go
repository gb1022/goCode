package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

/*代码会触发异常吗？请详细说明
答： 有可能触发异常，是随机事件。

解析：单个 chan 如果无缓冲时，将会阻塞。但结合 select 可以在多个 chan 间等待执行。有三点原则：

select 中只要有一个 case 能 return，则立刻执行。
当如果同一时间有多个 case 均能 return 则伪随机方式抽取任意一个执行。
如果没有一个 case 能 return 则可以执行“default”块。
此考题中的两个 case 中的两个 chan 均能 return，则会随机执行某个 case 块。故在执行程序时，有可能执行第二个 case，触发异常。
*/
