package main

import "fmt"

const (
	_ = iota
	a = 2022 + iota
	b
	c
	d
)

func main() {

	fmt.Println(a, b, c, d)

}