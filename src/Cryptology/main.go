package main

import (
	"fmt"
	utils "gin/src/Cryptology/Utils"
)

func main() {
	//base64测试
	str := "hello"                   //设置字符串的值
	ret := utils.Base64Encoding(str) //调用utils中的编码方法
	fmt.Println(ret)

	ret, err := utils.Base64Deoding(ret)
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println(ret)
	//测试URL编码解码
	str = "Good"
	ret = utils.Base64UrlEncoding(str)
	fmt.Println(ret)

	ret, err = utils.Base64UrlDecoding(ret) //解刚才编的码
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println(ret)

}
