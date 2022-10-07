package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //全部的元素
	c := a[3:]  //从4th开始到最后
	d := a[:6]  //前六个元素
	e := a[3:6] //从4th开始到6th结束
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
