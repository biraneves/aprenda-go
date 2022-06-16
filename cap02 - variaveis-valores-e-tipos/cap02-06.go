package main

import "fmt"

var y = "Olá! Bom dia!"		// Aqui, no escopo do pacote, não é possível o uso de :=

func main() {

	x := 10					// O operador curto de declaração só pode ser usado dentro de codeblocks
	z := 10.0

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	x = 20
	fmt.Printf("x: %v, %T\n\n", x, x)

	x, k := 30, 40
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("k: %v, %T", k, k)

}