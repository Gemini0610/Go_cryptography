package main

import (
	"fmt"
)

func main() { //slices 切片
	a := []int{1, 2, 3} //长度可以变化，是一个可以动态变化的数组
	fmt.Printf("lenth:%v", len(a))
	fmt.Printf("capcity%v", cap(a))

}
