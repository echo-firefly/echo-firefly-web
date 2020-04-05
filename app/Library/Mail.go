package Library

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(subject string, body string, mailTo []string, mailFrom string) error {
	mailConn := map[string]string{
		"host": "localhost",
		"port": "1025",
		"from": "localhost@wood.com",
		"name": mailFrom,
		"user": "",
		"pass": "",
	}

	port, err := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["name"]+"<"+mailConn["from"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                   //发送给多个用户
	m.SetHeader("Subject", subject)                                //设置邮件主题
	m.SetBody("text/plain", body)                                  //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err = d.DialAndSend(m)
	return err
}
