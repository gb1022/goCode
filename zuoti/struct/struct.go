package main

import (
	"fmt"
)

//Go 中没有继承！ 没有继承！没有继承！是叫组合！组合！组合！
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.People.ShowA()
	//	t.ShowB()
}
