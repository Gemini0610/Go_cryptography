package main
import(
	"fmt"
)
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'//字符的0
    //当我们直接输出byte值，就是输出了的对应的字符的码值
	//a 阿斯克码 97
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	//如果我们希望输出对应嘱咐，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c",c1,c2)

}