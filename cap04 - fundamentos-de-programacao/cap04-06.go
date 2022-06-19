package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d = iota
	e = iota
)

const (
	f = iota
	g
	_
	_
	h
)

func main() {

	fmt.Println(a, b, c, d, e)
	fmt.Println(f, g, h)

}