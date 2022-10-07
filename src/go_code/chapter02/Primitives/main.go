package main

import (
	"fmt"
)

func main() { //字符串
	s := "this is a string"
	s2 := "this is also a string"
	fmt.Printf("%v,%T\n", s+s2, s+s2) //用+号来将两个字符串进行合并

	//string UTF-8    rune UTF-32
	//string 用“”声明，rune用''声明
	r := 'a'
	fmt.Printf("%v,%T", r, r)
}
