package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	intChan <- 100
	intChan <- 200
	intChan <- 300
	intChan <- 500
	for i := 0; i < cap(intChan); i++ {
		i := <-intChan
		fmt.Println(i)
	}
}
