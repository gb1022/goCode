package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct {
}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Stduent{} //错误 无法编译
	//	var peo People = &Stduent{} //正确
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
