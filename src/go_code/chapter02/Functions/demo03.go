package main

import (
	"fmt"
)

func main() {
	greeting := "hello"
	name := "stacey"
	sayGreeting(&greeting, &name) //使用&获取地址
	fmt.Println(name)             //因为是对应地址，主函数变量随着其他函数改变而改变
}
func sayGreeting(greeting, name *string) { //指针*设置类型
	fmt.Println(*greeting, *name) //输出主函数中的变量
	*name = "Ted"                 //另一个函数修改name
	fmt.Println(*name)            //因为是对应地址
}
