package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota) //1024
	MB
	GB
)

func main() {
	fileSize := 400000000. //确保浮点数
	fmt.Printf("%.2fGB", fileSize/GB)
}
