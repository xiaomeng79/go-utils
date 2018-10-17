package random

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	n := 9
	s := GetRandomString(n)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Regexpf(t, "[0-9A-Za-z]{"+strconv.Itoa(n)+"}", s, "随机数格式不匹配")
}
