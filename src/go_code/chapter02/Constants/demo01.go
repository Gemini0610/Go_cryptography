package main

import (
	"fmt"
)

func main() {
	//Name Constants
	const myConst int = 42 //常量不能被修改
	//常量的类型：int float bool string
	fmt.Printf("%v,%T", myConst, myConst)
}
