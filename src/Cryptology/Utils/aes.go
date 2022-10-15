package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var key []byte = []byte("hallenhallenhall") //密钥

// 填充密码长度
func PadPwd(src_byte []byte, block_size int) []byte {

	// 16 13       13-3 = 10
	pad_num := block_size - len(src_byte)%block_size //计算需要填充的长度

	ret := bytes.Repeat([]byte{byte(pad_num)}, pad_num)

	src_byte = append(src_byte, ret...) //进行填充

	return src_byte

}

// 加密

func AesEncoding(src string) (string, error) {
	src_byte := []byte(src)
	// safer
	block, err := aes.NewCipher(key)
	if err != nil {
		return src, err
	}
	// 密码填充
	new_src_byte := PadPwd(src_byte, block.BlockSize())

	dst := make([]byte, len(new_src_byte)) //建立dst数组
	block.Encrypt(dst, new_src_byte)       //进行加密，将得到的放入dst中

	// base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)

	return pwd, nil

}

// 解密
func AesDecoding(pwd string) (string, error) {

	//pwd_byte := []byte(pwd)
	pwd_byte, err := base64.StdEncoding.DecodeString(pwd) //解码

	if err != nil {
		return pwd, err
	}

	block, err_block := aes.NewCipher(key)

	if err_block != nil {
		return pwd, err_block
	}

	dst := make([]byte, len(pwd_byte))
	block.Decrypt(dst, pwd_byte)
	// 填充的要去掉
	dst, _ = UnPadPwd(dst)

	return string(dst), nil

}

// 去掉填充的部分
func UnPadPwd(dst []byte) ([]byte, error) {

	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}

	// 去掉的长度
	unpad_num := int(dst[len(dst)-1])

	return dst[:(len(dst) - unpad_num)], nil

}
