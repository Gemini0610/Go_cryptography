package main

import (
	"fmt"
)

//const a int16 = 7

func main() { //主函数中的定义常量可以修改外面所定义的常量
	//不光可以修改值，还可以修改类型
	//const a int = 14
	const a = 14
	var b int16 = 27
	//fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", a+b, a+b) //相当于类型42+b
}
