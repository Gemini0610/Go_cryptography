package utils

import (
	"crypto/des"
	"encoding/base64"
)

// 支持24字节长度
var des3_key []byte = []byte("hellohellohellohellohell")

// 3des编码
func Des3Encoding(src string) (string, error) {
	//转为字节
	src_byte := []byte(src)
	//加密
	block, err := des.NewTripleDESCipher(des3_key) //用密钥加密
	if err != nil {
		return src, err
	}
	//自动填充
	new_src_byte := PadPwd(src_byte, block.BlockSize())
	//设置数组
	dst := make([]byte, len(new_src_byte))
	block.Encrypt(dst, new_src_byte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd, nil
}

// 3des解码
func Des3Decoding(pwd string) (string, error) {
	//base64解码
	pwd_byte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd, err
	}
	//解密转为字节
	block, err_block := des.NewTripleDESCipher(des3_key) //用密钥解密
	if err_block != nil {
		return pwd, err_block
	}
	dst := make([]byte, len(pwd_byte))
	block.Decrypt(dst, pwd_byte)
	//填充的去掉
	dst, _ = UnPadPwd(dst)
	return string(dst), nil
}
