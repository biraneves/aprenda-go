package main

import "fmt"

const (
	x = iota
	y
	z = iota << 1
)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {

	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Printf("%b\t%b\t%b\n", x, y, z)
	linha("=")
	fmt.Println("binário\t\t\t\t\t\tdecimal")
	linha("-")
	fmt.Printf("%b\t\t\t\t\t%d\n", KB, KB)
	fmt.Printf("%b\t\t\t\t%d\n", MB, MB)
	fmt.Printf("%b\t\t\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)
	linha("=")

}

func linha(char string) {
	linha := ""
	for i := 0; i < 64; i++ {
		linha += char
	}
	fmt.Println(linha)
}