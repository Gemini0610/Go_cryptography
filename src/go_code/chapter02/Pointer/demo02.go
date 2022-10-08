package main

import (
	"fmt"
)

func main() {
	var ms myStruct
	ms = myStruct{foo: 42}
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
