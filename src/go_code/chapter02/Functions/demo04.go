package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("This is", s)
}
func sum(values ...int) int {
	fmt.Println(values) //输出的是sum函数中的内容
	result := 0
	for _, v := range values {
		result += v //累加
	}
	return result //返回结果
}
