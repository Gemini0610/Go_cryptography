package utils

import "encoding/base64"

//base64编码
func Base64Encoding(str string) string {
	//设置参数str类型为string，返回值类型为string
	src := []byte(str)                            //将str转为src bype类型
	ret := base64.StdEncoding.EncodeToString(src) //括号中参数为src byte类型
	return ret
}

//base64解码
func Base64Deoding(str string) (string, error) {
	//DecodeString 将字符串转为字符以及错误的原因
	ret_byte, err := base64.StdEncoding.DecodeString(str)
	ret := string(ret_byte) //将解码的字符转为字符串
	return ret, err
}

//URL编码
func Base64UrlEncoding(str string) string {
	src := []byte(str)
	ret := base64.URLEncoding.EncodeToString(src)
	return ret
}

//URL解码
func Base64UrlDecoding(str string) (string, error) {
	ret_byte, err := base64.URLEncoding.DecodeString(str)
	ret := string(ret_byte)
	return ret, err
}
