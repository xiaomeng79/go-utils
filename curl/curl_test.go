package curl

import (
	"testing"
)

func TestCurl_Do(t *testing.T) {
	curl := New()
	curl.SetMethod("GET")
	curl.SetUrl("https://www.baidu.com/")
	curl.SetHeaders("Content-Type","application/json")
	err := curl.Do()
	if err != nil {
		t.Errorf("%v\n",err)
	}
}
