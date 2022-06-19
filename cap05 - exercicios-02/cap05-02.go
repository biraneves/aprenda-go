package main

import "fmt"

const a int = 20
const b int = 30

func main() {

	c := a == b
	d := a <= b
	e := a < b
	f := a >= b
	g := a > b
	h := a != b

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	linha("=")
	fmt.Printf("a == b: %v\n", c)
	fmt.Printf("a <= b: %v\n", d)
	fmt.Printf("a < b: %v\n", e)
	fmt.Printf("a >= b: %v\n", f)
	fmt.Printf("a > b: %v\n", g)
	fmt.Printf("a != b: %v\n", h)

}

func linha(char string) {
	linha := ""
	for i := 0; i < 64; i++ {
		linha += char
	}
	fmt.Println(linha)
}