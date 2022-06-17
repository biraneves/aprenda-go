package main

import "fmt"

type hotdog int

var a hotdog = 10
var b int = 10

func main() {

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	b = int(a) 					// Conversão de tipo

}