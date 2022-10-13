package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 330
	intChan <- 20
	close(intChan)
	//Channel关闭后不能再写入数据，但是可以读取数据
	n1 := <-intChan
	fmt.Println("n1=", n1)
}
