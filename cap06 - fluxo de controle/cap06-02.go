package main

import "fmt"

func main() {

	for horas := 0; horas < 12; horas++ {

		fmt.Println("Hora:", horas)

		for minutos := 0; minutos < 60; minutos++ {
			
			fmt.Print(minutos, " ")

		}

		fmt.Println("")
		linha("-")

	}

}

func linha(char string) {

	linha := ""
	for i := 0; i < 64; i++ {
		linha += char
	}
	fmt.Println(linha)

}