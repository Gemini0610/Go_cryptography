package main

import (
	"fmt"
)

func main() {
	//演示管道的作用
	var intChan chan int        //声明
	intChan = make(chan int, 3) //初始化 3是容量

	//fmt.Printf("intchan的值=%v\n", intChan)
	//向管道写入数据

	intChan <- 10
	num := 100
	intChan <- num
	//写入数据不能超过管道的定义的容量
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
}
