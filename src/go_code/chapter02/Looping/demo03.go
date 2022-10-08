package main

import (
	"fmt"
)

func main() {
	i := 0
	for ; i < 5; i++ { //之前定义的变量可以在for语句中省略
		fmt.Println(i)
	}
}
