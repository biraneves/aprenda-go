package main

import "fmt"

var x bool

func main() {

	fmt.Println(x)
	x = true
	fmt.Println(x)

	x = 10 < 100
	y := 10 == 100
	z := 10 > 100

	fmt.Println(x, y, z)

}