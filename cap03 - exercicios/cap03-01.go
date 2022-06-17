package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	linha()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	linha()
	fmt.Printf("x: %v - y: %v - z: %v\n", x, y, z)

}

func linha() {
	fmt.Println("--------------------------------------------------")
}