package main
import(
	"fmt"
	"unsafe"
)
//定义全局变量

func main() {
	var i =100
	
	//如何查看某个变量的数据类型
	fmt.Printf("i的类型", i)
    
	//如何在程序查看某个变量的占用字节大小和数据类型（使用较多）
	var n2 int64 =10
    fmt.Printf("n2的数据类型%T,n2占用%d",n2, unsafe.Sizeof(n2))
}
