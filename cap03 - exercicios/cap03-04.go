package main

import "fmt"

type nika int
var x nika

func main() {

	fmt.Printf("x: %v - %T\n", x, x)

	x = 42

	fmt.Println("x:", x)

}