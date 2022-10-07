package main
import "fmt"
func main() {
	//1、指定类型变量，声明后若不赋值，使用默认值
	var i int//int 默认值为0
	fmt.Println("i=",i)

	//2、根据值自行判定变量类型（类型推导）
	var num = 10.11
	fmt.Println("num =",num)

	//第三种：省略var， :=左侧的变量不应该是已经声明过的
	name :="tom"
	//等价于 var name string    name="tom"
	fmt.Println(name)
}