package main

import (
	"fmt"
)

func main() {
	greeting := "hello"
	name := "stacey"
	sayGreeting(greeting, name)
	fmt.Println(name) //主函数的变量不随着其他函数的变量改变而改变
}
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name) //先输出主函数定义的变量
	name = "Ted"                //另一个函数修改name
	fmt.Println(name)           //将修改的name输出
}
