//随机数
package random

import (
	"fmt"
	"math/rand"
	"time"
)

//获取n位随机字符串
func GetRandomString(leng int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
返回8位随机数
*/
func EightBitRand() string {
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", _rand.Int31n(100000000))
}
