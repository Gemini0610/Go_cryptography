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

	//AES编码
	pwd_Aes, _ := utils.AesEncoding("hall")
	fmt.Println("AES编码:", pwd_Aes)
	ret_Aes, _ := utils.AesDecoding(pwd_Aes)
	fmt.Println("AES译码:", ret_Aes)

	//Des编码
	pwd_Des, _ := utils.DesEncoding("dalian")
	fmt.Println("DES编码:", pwd_Des)
	ret_Des, _ := utils.DesDecoding(pwd_Des)
	fmt.Println("DES译码:", ret_Des)
	//3Des编码
	pwd_3Des, _ := utils.Des3Encoding("dalian")
	fmt.Println("3DES编码:", pwd_3Des)
	ret_3Des, _ := utils.Des3Decoding(pwd_3Des)
	fmt.Println("3DES译码:", ret_3Des)

	//RSA加密
	ret_RSA, _ := utils.RsaEncoding("hallen", "publicKey.pem")

	fmt.Println("RSA编码:", ret_RSA)
	//RSA解密
	ret_pri, _ := utils.RsaDecoding(ret_RSA, "privateKey.pem")

	fmt.Println("RSA解码:", string(ret_pri))

	//数字签名
	sign, err_sign := utils.RsaGetSign("Dalian", "privateKey.pem")
	if err_sign != nil {
		fmt.Println("出错:", err_sign)
		panic("签名失败")
	}
	fmt.Println("签名为:", sign)
	ok, err_varify := utils.RsaVarifySign(sign, "Dalian", "public.pem")
	if err_varify != nil {
		fmt.Println(err_varify)
		panic("验证失败")
	}
	fmt.Println("验证:", ok)
}
