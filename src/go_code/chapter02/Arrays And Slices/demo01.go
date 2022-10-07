package main

import (
	"fmt"
)

func main() { //create Arrays
	//grades := [...]int{97, 85, 93}//[...]可以自动设置大小
	var students [6]string
	fmt.Printf("Students:%v\n", students)
	students[1] = "Lisa"
	students[2] = "Sam"
	students[0] = "Amy"
	fmt.Printf("Students:%v\n", students)
	fmt.Printf("Students#1:%v\n", students[1])

	//内置函数
	//len()求数组的长度
	fmt.Printf("Number of Students:%v\n", len(students))
}
