package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenSha256(src string) string {
	src_byte := []byte(src) //字符串转成字符数组
	sha256_new := sha256.New()
	sha256_bytes := sha256_new.Sum(src_byte)
	sha256_string := hex.EncodeToString(sha256_bytes[:])
	return sha256_string
}
