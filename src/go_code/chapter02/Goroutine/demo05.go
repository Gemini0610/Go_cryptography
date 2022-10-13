package main

import (
	"fmt"
	"runtime"
)

func main() {
	//get CPU number
	num := runtime.NumCPU()
	//set num-1 cpu go program
	runtime.GOMAXPROCS(num)
	fmt.Println("num", num)
}
