package main

import (
	"fmt"
)

func main() {
Loop: //标签
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 4 {
				break Loop
			}
		}
	}
}
