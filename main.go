package main

import (
	"fmt"
)

const (
	a = iota
	b
	c = 5
	d
	e = iota
)

func main() {
	fmt.Println(a, b, c, d, e)
}
