package email

import (
	"net/smtp"
	"strings"
)

const (
	EMAIL_HOST        = "smtp.exmail.qq.com" // 使用腾讯的邮箱
	EMAIL_SERVER_ADDR = "smtp.exmail.qq.com:25"
	EMAIL_USER        = "xxx@qq.com" //发送邮件的邮箱
	EMAIL_PASSWORD    = "xxxxx"      //邮箱密码
)

type Email struct {
	Host     string
	Server   string
	Account     string
	Password string
}

func (e *Email) SendMail(subject, content string, sendTo []string) error {
	auth := smtp.PlainAuth("", e.Account, e.Password, e.Host)
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		contentType := "Content-Type: text/plain" + "; charset=UTF-8"
		for _, v := range sendTo {
			str := strings.Replace("From: "+e.Account+"~To: "+v+"~Subject: "+subject+"~"+contentType+"~~", "~", "\r\n", -1) + content
			err := smtp.SendMail(
				e.Server,
				auth,
				e.Account,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()
	// 同步
	for i := 0; i < len(sendTo); i++ {
		// 可以记录error日志
		<-done
	}

	return nil
}

func NewEmail() *Email {
	return &Email{
		Host: EMAIL_HOST,
		Server: EMAIL_SERVER_ADDR,
		Account: EMAIL_USER,
		Password: EMAIL_PASSWORD,
	}
}