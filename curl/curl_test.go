package curl

import (
	"testing"
)

func TestCurl_Do(t *testing.T) {
	//以下设置不不要的可以不设置
	curl := New()                                      //创建对象
	curl.SetMethod("GET")                              //设置请求方法
	curl.SetUrl("https://www.baidu.com/")              //设置Url
	curl.SetHeader("Content-Type", "application/json") //设置请求类型
	curl.SetBody("")                                   //设置请求体，
	curl.AddHeader("test", "test01")                   //增加请求头test的值
	curl.AddHeader("test", "test02")                   //增加请求头
	err := curl.Do()
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestCurlBuilder_Build(t *testing.T) {
	//curl也可以通过链式构建,
	cb := &CurlBuilder{}
	curl := cb.SetMethod("GET").SetUrl("https://www.baidu.com/").SetHeader("Content-Type", "application/json").Build()
	//执行请求
	err := curl.Do()
	if err != nil {
		t.Errorf("%v\n", err)
	}
}
