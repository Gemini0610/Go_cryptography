package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("no one or two")
	}
}
