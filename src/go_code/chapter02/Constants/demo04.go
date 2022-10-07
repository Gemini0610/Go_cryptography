package main

import (
	"fmt"
)

const (
	catSpecialist    = iota //默认为0
	dogSpecialist           //叠加为1
	snakeSpercialist        //叠加为2
)

func main() {
	var specialistType int                              //默认值为0
	fmt.Printf("%v\n", specialistType == catSpecialist) //判断值
}
