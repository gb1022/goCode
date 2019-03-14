package main

import (
	"fmt"
	//	"os"
	"time"
)

/*有四个线程1、2、3、4。线程1的功能就是输出1，线程2的功能就是输出2，以此类推.........现在有四个文件ABCD。初始都为空。现要让四个文件呈如下格式：

A：1 2 3 4 1 2....

B：2 3 4 1 2 3....

C：3 4 1 2 3 4....

D：4 1 2 3 4 1....
*/
var s1 = make(chan int, 4)
var s2 = make(chan int, 4)
var s3 = make(chan int, 4)
var s4 = make(chan int, 4)

//我这里用四个切片搞定
var sa []int
var sb []int
var sc []int
var sd []int

func f1() {
	s1 <- 1
}
func f2() {
	s2 <- 2
}
func f3() {
	s3 <- 3
}
func f4() {
	s4 <- 4
}

func writesa() {
	//	fmt.Println("2222222222222")
	i1 := <-s1
	//	fmt.Println("3333333333333,i1:", i1)
	i2 := <-s2
	//	fmt.Println("4444444444444")
	i3 := <-s3
	//	fmt.Println("5555555555555555")
	i4 := <-s4
	//	fmt.Println("66666666666666")
	sa = append(sa, i1)
	//	fmt.Println("777777777777777777")
	sa = append(sa, i2)
	sa = append(sa, i3)
	sa = append(sa, i4)
}
func writesb() {
	i1 := <-s1
	i2 := <-s2
	i3 := <-s3
	i4 := <-s4
	sb = append(sb, i2)
	sb = append(sb, i3)
	sb = append(sb, i4)
	sb = append(sb, i1)
}
func writesc() {
	i1 := <-s1
	i2 := <-s2
	i3 := <-s3
	i4 := <-s4
	sc = append(sc, i3)
	sc = append(sc, i4)
	sc = append(sc, i1)
	sc = append(sc, i2)
}
func writesd() {
	i1 := <-s1
	i2 := <-s2
	i3 := <-s3
	i4 := <-s4
	sd = append(sd, i4)
	sd = append(sd, i1)
	sd = append(sd, i2)
	sd = append(sd, i3)
}

func main() {
	for s := 0; s < 100; s++ {
		go f1()
		go f2()
		go f3()
		go f4()
	}
	time.Sleep(time.Second * 2)
	for i := 0; i < 100; i++ {
		//		fmt.Println("1111111111111")
		go writesa()
		go writesb()
		go writesc()
		go writesd()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("len(sa):", len(sa), "A:", sa)
	fmt.Println("len(sb):", len(sb), "B:", sb)
	fmt.Println("len(sc):", len(sc), "C:", sc)
	fmt.Println("len(sd):", len(sd), "D:", sd)

}
