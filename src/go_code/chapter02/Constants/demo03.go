package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)
const (
	a2 = iota
)

func main() { //常量块中相同的含义自动+1，不同常量块互不影响
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
}
