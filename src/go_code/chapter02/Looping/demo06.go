package main

import (
	"fmt"
)

func main() {
	s := [3]int{1, 2, 3}  //array
	for k, v := range s { //遍历arrays
		fmt.Println(k, v)
	}

	a := "hello world"

	for k, v := range a {
		fmt.Println(k, string(v)) //学会转换
	}
}
