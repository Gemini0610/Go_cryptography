package main

import (
	"fmt"
)

func main() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 1) //添加元素
	fmt.Println(a)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 2, 3, 4, 5, 7, 10, 11, 12, 13) //添加元素
	fmt.Println(a)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}
