package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := "a"
	b := "á"
	c := "的"

	fmt.Println(a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Println(d, e, f)

	fmt.Println()
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	
}