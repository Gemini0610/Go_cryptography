package main

import (
	"fmt"
)

func main() { //多个元素进行for statement
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}
