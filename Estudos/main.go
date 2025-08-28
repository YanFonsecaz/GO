package main

import (
	"fmt"
	"estudos/internal/pacote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(pacote.Foo)
	fmt.Println(somar(1,2))

	c,d := swap(10,20)
	fmt.Println(c,d)
}

func somar (a,b int) int {
	return a + b
}

func swap (c,d int) (int,int) {
	return d,c
}