package main

import "fmt"

const (
	x = 10
	y = 20
	z = 30
)

// var y = 10
// var a int
var b float64

func main() {

	b = x
	fmt.Printf("%T, %T\n", b, x)
	fmt.Println(x, y, z)

}