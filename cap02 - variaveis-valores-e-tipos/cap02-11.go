package main

import "fmt"

type hotdog int			// Criação de um tipo personalizado
var b hotdog

func main() {

	fmt.Printf("%T\n", b)

}