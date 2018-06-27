package httpclient

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	//测试是否是单例
	assert.Exactly(t,New(),New(),"不是单例模式")
}
