package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 保存生成的公钥密钥
func SaveRsaSignKey(bits int) error {
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

// 获取公钥
func GetPublicKey(file_path string) (*rsa.PublicKey, error) { //公钥加密
	//打开文件
	file, err := os.Open(file_path)
	if err != nil {
		return &rsa.PublicKey{}, err
	}

	//获取文件的信息
	file_info, err_info := file.Stat()
	if err_info != nil {
		return &rsa.PublicKey{}, err_info
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
		return &rsa.PublicKey{}, err_pb
	}
	return publicKey, nil
}

// 获取私钥
func GetPrivateKey(file_path string) (*rsa.PrivateKey, error) { //私钥解密
	//获取文件
	file, err := os.Open(file_path)
	if err != nil {
		return &rsa.PrivateKey{}, err
	}
	//获取文件信息
	file_info, err_info := file.Stat()
	if err_info != nil {
		return &rsa.PrivateKey{}, err
	}
	//读取内容
	//建立数组
	key_bytes := make([]byte, file_info.Size())
	file.Read(key_bytes)
	//pem解码
	pem.Decode(key_bytes)
	//x509解码
	privateKey, err_pb := x509.ParsePKCS1PrivateKey(key_bytes)
	if err_pb != nil {
		return &rsa.PrivateKey{}, err_pb
	}
	return privateKey, nil
}

// 数字签名
// 私钥签名，公钥验证签名
func RsaGetSign(src, file_path string) ([]byte, error) {
	//获取私钥
	privateKey, err := GetPrivateKey(file_path)
	if err != nil {
		return []byte{}, err
	}
	sha_new := sha256.New()
	src_byte := []byte(src)
	sha_bytes := sha_new.Sum(src_byte)
	//签名 hash用cryto有sha256
	sign, err_sign := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sha_bytes)
	if err_sign != nil {
		return []byte{}, err_sign
	}
	return sign, nil
}

// 验证签名使用公钥
// sign签名, src明文, filepath获取公钥路径  返回布尔类型和错误
func RsaVarifySign(sign []byte, src, file_path string) (bool, error) {
	publicKey, err := GetPublicKey(file_path)
	if err != nil {
		return false, err
	}
	//hash
	sha_new := sha256.New()
	src_byte := []byte(src)
	sha_new.Write(src_byte)
	sha_bytes := sha_new.Sum(src_byte)
	//验证签名
	err_varify := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, sha_bytes, sign)
	if err_varify != nil {
		return false, err_varify
	}
	return true, nil
}
