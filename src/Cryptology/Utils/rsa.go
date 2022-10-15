package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 保存生成的公钥密钥
func SaveRsaKey(bits int) error {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		fmt.Println(err)
		return err
	}
	//通过私钥创建公钥
	publicKey := privateKey.PublicKey
	//使用x509标准进行编码,生成ASN.1编码字符串
	x509_Private := x509.MarshalPKCS1PrivateKey(privateKey)
	x509_Public := x509.MarshalPKCS1PublicKey(&publicKey)
	//封装block结构数据
	block_Private := pem.Block{Type: "private key", Bytes: x509_Private}
	block_Public := pem.Block{Type: "public key", Bytes: x509_Public}
	//创建存放公私钥的文件
	privateFile, err_pri := os.Create("privateKey.pem")
	if err_pri != nil {
		return err_pri
	}
	defer privateFile.Close()
	pem.Encode(privateFile, &block_Private)
	publicFile, err_pub := os.Create("pulicKey.pem")
	if err_pub != nil {
		return err_pub
	}
	defer publicFile.Close()
	pem.Encode(publicFile, &block_Public)
	return nil
}

// 加密   函数(明文，公钥路径)
func RsaEncoding(src, file_path string) ([]byte, error) {
	//明文转为数组byte
	src_byte := []byte(src)
	//打开文件
	file, err := os.Open(file_path)
	if err != nil {
		return src_byte, err
	}

	//获取文件的信息
	file_info, err_info := file.Stat()
	if err_info != nil {
		return src_byte, err_info
	}

	//读取文件内容
	//设置数组
	key_bytes := make([]byte, file_info.Size())
	//将读取的内容放入到数组中
	file.Read(key_bytes)

	//文件解码
	//pem解码
	block, _ := pem.Decode(key_bytes)
	//x509解码
	publicKey, err_pb := x509.ParsePKCS1PublicKey(block.Bytes)
	if err_pb != nil {
		return src_byte, err_pb
	}
	//使用公钥对明文进行加密

	ret_byte, err_rb := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src_byte)
	if err_rb != nil {
		return ret_byte, err_rb
	}
	fmt.Println(ret_byte)
	return ret_byte, nil
}

// 解密   函数(密码数组,私钥路径)
func RsaDecoding(src_byte []byte, file_path string) ([]byte, error) {
	//打开文件
	file, err := os.Open(file_path)
	if err != nil {
		return src_byte, err
	}
	//获取文件信息
	file_info, err_info := file.Stat()
	if err_info != nil {
		return src_byte, err
	}
	//读取文件内容
	//建立数组 根据文件的信息长度建立
	key_bytes := make([]byte, file_info.Size())
	//将文件内容放入数组中
	file.Read(key_bytes)

	//pem解码
	pem.Decode(key_bytes)
	//x509解码
	privateKey, err_pb := x509.ParsePKCS1PrivateKey(key_bytes)
	if err_pb != nil {
		return src_byte, err_pb
	}
	//进行解密
	ret_byte, err_ret := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src_byte)
	if err_ret != nil {
		return ret_byte, err_ret
	}
	return ret_byte, nil
}
