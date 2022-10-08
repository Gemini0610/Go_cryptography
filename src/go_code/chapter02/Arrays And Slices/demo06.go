package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3, 100) //通过make 来创建切片,make也会创建一个数组，是由切片在底层维护
	fmt.Println(a)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
