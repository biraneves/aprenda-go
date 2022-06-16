package main

import "fmt"

func main() {
	_, erro := fmt.Println("Hello, World!", "Another string", "Just one more!")
	fmt.Println("Erro:", erro)
}