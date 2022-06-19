package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)
	linha()
	fmt.Println("Os valores acima chamam-se zeros.")

}

func linha() {
	fmt.Println("--------------------------------------------------")
}