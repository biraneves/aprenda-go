package main

import "fmt"

func main() {

	x := 0

	for x < 10 {

		fmt.Printf("%v: Xis é menor que dez\n", x)
		x++

	}

	linha("-")

	x = 0

	for {

		if x < 10 {

			fmt.Printf("%v: xis é menor do que dez\n", x)
			x++

		} else {
			fmt.Printf("%v: xis é maior ou igual a dez\n", x)
			break
		}

	}

}

func linha(char string) {

	linha := ""
	for x := 0; x < 64; x++ {
		linha += char
	}
	fmt.Println(linha)

}