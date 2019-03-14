package main

import (
	"fmt"
)

func main() {
	a := make([]byte, 2, 5)
	a1 := append(a, 'a')
	a2 := append(a, 'b')

	fmt.Printf("a1=====> %s %p\n", a1, a1)
	fmt.Println("a1:", a1, len(a2))
	fmt.Printf("a2=====> %s %p\n", a2, a2)
	fmt.Println("a2:", a2, len(a2))

	a2[1] = 'c'
	fmt.Printf("a=====> %s %p\n", a, a)
	fmt.Println("a:", a, len(a2))

}
