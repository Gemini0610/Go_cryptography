package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量 type 类型 值
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)
	//将rvalue转成interface
	iV := rValue.Interface()
	//将interface()通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

// 对结构体的反射
func reflectTest02(b interface{}) {
	rValue := reflect.ValueOf(b)
	iV := rValue.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//定义一个int
	var num int = 100
	reflectTest01(num)

	//定义学生实例
	stu := Student{
		Name: "Tom",
		Age:  10,
	}
	reflectTest02(stu)
}
