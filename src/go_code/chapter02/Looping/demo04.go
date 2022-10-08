package main

import (
	"fmt"
)

func main() { //极简for
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}
