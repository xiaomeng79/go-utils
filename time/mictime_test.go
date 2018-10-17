package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//测试思路
//1.生成毫秒时间戳，查看长度是否正确
//2.毫秒时间戳转字符串
//3.字符串转毫秒时间
var (
	timeValues = []struct {
		key string
		mic string
		str string
	}{
		{"k1", "1528879610000", "2018-06-13 16:46:50"},
		{"k2", "1528879610000", "2018/06/13 16:46:50"},
		{"k3", "1523021ssssss", "2018-06-13 16:46:50"},
		{"k4", "1528879610000", "aaaa-bb-cc dd:ee:ff"},
	}
)

//生成时间戳
func TestGenMicTime(t *testing.T) {
	micTime := GenMicTime()
	assert.Lenf(t, micTime, 13, "生成时间格式长度不正确:%d", len(micTime))
}

//时间戳转字符串
func TestMicTimeToStr(t *testing.T) {
	for _, v := range timeValues {
		res, err := MicTimeToStr(v.mic)
		if v.key == "k1" {
			assert.Equalf(t, v.str, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else if v.key == "k2" || v.key == "k4" {
			assert.NotEqualf(t, v.str, res, "key:%s", v.key)
			assert.NoErrorf(t, err, "key:%s", v.key)
		} else {
			assert.Errorf(t, err, "key:%s", v.key)
		}

	}
}

//字符串转时间戳
func TestStrToMicTime(t *testing.T) {
	for _, v := range timeValues {
		res := StrToMicTime(v.str)
		if v.key == "k1" {
			assert.Equalf(t, v.mic, res, "key:%s", v.key)
		} else {
			assert.NotEqualf(t, v.mic, res, "key:%s", v.key)
		}

	}
}
