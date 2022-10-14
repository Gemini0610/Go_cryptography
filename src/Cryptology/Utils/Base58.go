package utils

import (
	"bytes"
	"fmt"
	"math/big"
)

var b58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// base58区块链专用
func Base58Encoding(src string) string {
	//he (Asc码)编码过程: 104 101 -->104*256+101=26725
	// (若有第三个 则101也要乘256)
	//解码过程:26725/58 余数 在编码表中找对应的值

	//asc码对应值
	src_byte := []byte(src)
	fmt.Println("编码为:", src_byte)

	//转成十进制
	i := big.NewInt(0).SetBytes(src_byte)
	fmt.Println("将编码转为十进制:", i)
	var mod_slice []byte //容器 指针
	//循环取余数
	for i.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0)
		i58 := big.NewInt(58)
		i.DivMod(i, i58, mod)
		mod_slice = append(mod_slice, b58[mod.Int64()]) //添加到容器中
	}
	fmt.Println(mod_slice)
	//遇到0使用字节1代替
	for _, s := range src_byte {
		if s != 0 {
			break
		}
		mod_slice = append(mod_slice, byte('1'))
	}
	fmt.Println(mod_slice)

	ret_mode_slice := ReverseByteArr(mod_slice) //调转
	return string(ret_mode_slice)               //返回string类型
}

// byte数组进行反转
// 反转方式1
func ReverseByteArr(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i] //调转
	}
	return b
}

// // 反转方式2
// func ReverseByteArr2(b []byte) []byte { //0-9 从0开始所以len长度-1
// 	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
// 		//i是左边指针为0,j为右边指针,i<j
// 		b[i], b[j] = b[j], b[i] //j的值赋给i，i的值赋给j
// 	}
// }

func Base58Decoding(src string) string {
	src_byte := []byte(src) //转为byte数组
	ret := big.NewInt(0)    //得到十进制
	for _, b := range src_byte {
		//获取index
		i := bytes.IndexByte(b58, b)
		fmt.Println("获取index为:", i)
		ret.Mul(ret, big.NewInt(58))       //乘法
		ret.Add(ret, big.NewInt(int64(i))) //加法
	}
	return string(ret.Bytes())
}
