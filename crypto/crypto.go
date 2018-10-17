//封装一下常用的加密的类型
package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

//MD5加密
func MD5(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

//SHA1加密
func SHA1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

//SHA256加密
func SHA256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

//SHA512加密
func SHA512(s string) string {
	r := sha512.Sum512([]byte(s))
	return hex.EncodeToString(r[:])
}
