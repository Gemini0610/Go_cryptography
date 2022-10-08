package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("somthing bad happened")
	fmt.Println("end")
}
