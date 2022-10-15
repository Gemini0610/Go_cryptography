package utils

import (
	"crypto/des"
	"encoding/base64"
)

// 密钥8字节
var des_key []byte = []byte("hellohhh")

// 加密
func DesEncoding(src string) (string, error) {
	src_byte := []byte(src)

	block, err := des.NewCipher(des_key)
	if err != nil {
		return src, err
	}
	//密码填充
	new_src_byte := PadPwd(src_byte, block.BlockSize())
	//设置数组
	dst := make([]byte, len(new_src_byte))
	//加密后放入数组中
	block.Encrypt(dst, new_src_byte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)

	return pwd, nil
}

// 解密
func DesDecoding(pwd string) (string, error) {
	pwd_byte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd, err
	}
	block, err_block := des.NewCipher(des_key)
	if err_block != nil {
		return pwd, err_block
	}
	dst := make([]byte, len(pwd_byte))
	block.Decrypt(dst, pwd_byte) //解密
	//去掉填充的
	dst, _ = UnPadPwd(dst)
	return string(dst), nil

}
