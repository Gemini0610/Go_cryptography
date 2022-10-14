package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(src string) string {
	src_byte := []byte(src)
	//md5计算
	md5_new := md5.New()
	md5Bytes := md5_new.Sum(src_byte)
	md5String := hex.EncodeToString(md5Bytes)
	return md5String
}
