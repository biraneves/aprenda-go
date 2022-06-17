package main

import "fmt"

type nika int
var x nika
var y int

func main() {

	fmt.Printf("x: %v - %T\n", x, x)
	x = 42
	fmt.Println("x:", x)

	y = int(x)
	fmt.Printf("y: %v = %T\n", y, y)

}