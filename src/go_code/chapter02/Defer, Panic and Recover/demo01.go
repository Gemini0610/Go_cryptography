package main

import (
	"fmt"
)

func main() {
	a := "start"
	defer fmt.Println(a) //虽然a改变了值，但这个程序是最后运行的
	a = "end"
}
