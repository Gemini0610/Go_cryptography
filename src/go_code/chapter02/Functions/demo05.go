package main

import (
	"fmt"
)

// 斐波那契数列
func fbn(value int) int {
	if value == 1 || value == 2 {
		return 1
	} else {
		return fbn(value-1) + fbn(value-2)
	}
}
func main() {
	for n := 1; n < 10; n++ {
		fmt.Println(fbn(n))
	}

}
