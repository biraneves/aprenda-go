package main

import "fmt"

func main() {

	x := 10					// O operador curto de declaração é usado na declaração de uma nova variável
	y := "Olá! Bom dia!"
	z := 10.0

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	x = 20					// Este é o operador de atribuição
	fmt.Printf("x: %v, %T\n\n", x, x)

	x, k := 30, 40			// Aqui é possível o uso de := porque uma das variáveis é nova
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("k: %v, %T", k, k)

}