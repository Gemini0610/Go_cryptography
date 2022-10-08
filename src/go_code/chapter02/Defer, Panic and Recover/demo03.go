package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  //说明捕获到错误
			log.Println("Error:", err) //输出
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
