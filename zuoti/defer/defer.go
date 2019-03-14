package main

import (
	"fmt"
	"time"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	time.Sleep(time.Second)
}

/*
输出结果为：

1
2
3
4
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
解析：在解题前需要明确两个概念： + defer 是在函数末尾的 return 前执行，先进后执行，具体见问题1。 + 函数调用时 int 参数发生值拷贝。

不管代码顺序如何，defer calc func 中参数 b 必须先计算，故会在运行到第三行时，执行 calc("10",a,b) 输出：10 1 2 3 得到值 3，将 cal("1",1,3) 存放到延后执执行函数队列中。

执行到第五行时，现行计算 calc("20", a, b) 即 calc("20", 0, 2) 输出：20 0 2 2 得到值 2，将 cal("2",0,2) 存放到延后执行函数队列中。

执行到末尾行，按队列先进后出原则依次执行：cal("2",0,2)、cal("1",1,3) ，依次输出：2 0 2 2、1 1 3 4
*/
