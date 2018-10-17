//加密发送邮件
package email

import (
	"fmt"
)

func ExampleSendEmailTls() {
	//发送方
	host := "smtp.mxhichina.com" //使用的邮件服务器
	port := "465"                //ssl端口
	username := "aa@tx.com"      //发送方邮箱
	nickname := "test"           //发送方昵称
	password := "test"           //密码

	//接收方
	sendTo := "abc@dd.com"
	subject := "发送主题"
	body := "发送的内容,可以是html格式,<p>hello world</p>"

	err := SendEmailTls(host, port, username, nickname, password, sendTo, subject, body)
	//填写参数正确返回<nil>,不正确返回如下错误信息
	fmt.Println(err)
	//Output:
	//526 Authentication failure[0]

}
