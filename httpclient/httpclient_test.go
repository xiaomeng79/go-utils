package httpclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	//测试是否是单例
	assert.Exactly(t, New(), New(), "不是单例模式")
}
