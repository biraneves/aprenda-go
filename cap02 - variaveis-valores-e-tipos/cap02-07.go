package main

import "fmt"

var y = 10		// var é necessário para declarar uma variável em escopo de pacote

func main() {
	qualquerCoisa(y)
}

func qualquerCoisa (x int) {
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}