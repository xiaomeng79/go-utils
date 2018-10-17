//加密发送邮件
package email

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"strings"
)

//host:邮件服务器地址
//port:邮件服务器加密端口
//password：邮件服务器密码
//username：邮件服务器用户名
//nickname:发送方昵称
//sendTo：接收者，多个逗号隔开
//subject：发送主题
//body:发送内容
func SendEmailTls(host, port, username, nickname, password, sendTo, subject, body string) error {
	auth := smtp.PlainAuth("", username, password, host)
	to := strings.Split(strings.Replace(sendTo, "，", ",", -1), ",")
	user := username
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

	_port, _ := strconv.Atoi(port)
	return sendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, _port),
		auth,
		user,
		to,
		msg,
	)
}

//smtp客户端
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
