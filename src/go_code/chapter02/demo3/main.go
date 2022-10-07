package main
import "fmt"
//定义全局变量
var(
	n6=300
	name2="tom"
)
func main() {

	//golang如何一次性声明多个变量方式1
	var n1,n2,n3 int
	fmt.Println(n1,n2,n3)

	//一次性声明多个变量的方式2
	//var n4,name,n5 = 100,"tom",888.5
	//fmt.Println(n4,name,n5)

    //3、同样可以使用类型推导
	n4,name,n5 := 100,"tom",88.4

	fmt.Println(n4,name,n5)
	fmt.Println(name2,n6)
	fmt.Printf("数据类型",name2)

}