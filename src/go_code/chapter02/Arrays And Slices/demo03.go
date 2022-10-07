package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	b := &a  //令b=a,加上&后 a b 的值一起修改
	b[1] = 5 //修改第二个元素的值
	fmt.Println(a)
	fmt.Println(b)
}
