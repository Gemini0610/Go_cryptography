package main
import(
	"fmt"

)
//定义全局变量

func main() {
	//浮点数尾数可能有精度损失
	var num3 float32 = -123.00000901
	var num4 float64 = -123.00000901
	fmt.Println("num3=",num3,"num4=",num4)

	//科学计数法格式
	num5 := 5.124e2
	fmt.Println("num5=",num5)
}
