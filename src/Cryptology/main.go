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

	//base58编码
	ret_b58 := utils.Base58Encoding("he")
	fmt.Println(ret_b58) //输出调转后的编码
	// ret_byte := utils.ReverseByteArr([]byte{110, 119, 56})
	// fmt.Println(ret_byte)
	aa := utils.Base58Decoding("8wn")
	fmt.Println("Base58解码为:", aa)

	//md5编码
	ret_md5 := utils.GenMd5("Sam")
	fmt.Println("md5编码为:", ret_md5)
	//sha256编码
	ret_sha256 := utils.GenSha256("Sam")
	fmt.Println("sha256编码为:", ret_sha256)
}
