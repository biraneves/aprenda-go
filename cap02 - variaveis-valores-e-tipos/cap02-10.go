package main

import "fmt"

func main() {

	x := "Olá, bom dia!\nComo vai?\tEspero que \"tudo bem\"."		// Interpreted string literal
	y := `"Olá, bom dia!\nComo vai?\tEspero que \"tudo bem\"."`		// Raw string literal
	
	fmt.Println(x)
	fmt.Println(y)

	a := "Oi!"
	b := "Bom dia!"

	c := fmt.Sprint(a, b)

	fmt.Println(c)

	i := 5
	j := 8

	k := fmt.Sprint(i, j)

	fmt.Println(k)

}