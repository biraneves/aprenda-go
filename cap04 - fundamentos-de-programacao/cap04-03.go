package main

import "fmt"

func main() {

	s := "ASCII, áéíçãò, 的"
	sb := []byte(s)
	
	fmt.Printf("%v\t%T\n", s, s)
	fmt.Printf("%v\t%T\n", sb, sb)

	linha()

	for _, v := range sb {

		fmt.Printf("%b\t%v\t%T\t%#U\t%#x\n", v, v, v, v, v)

	}

	linha()

	for _, v := range s {

		fmt.Printf("%b\t%v\t%T\t%#U\t%#x\n", v, v, v, v, v)

	}

	linha()

	for i := 0; i < len(s); i++ {

		fmt.Printf("%b\t%v\t%T\t%#U\t%#x\n", s[i], s[i], s[i], s[i], s[i])

	}

}

func linha() {
	fmt.Println("--------------------------------------------")
}