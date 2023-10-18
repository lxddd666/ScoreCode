package sys

import (
	"fmt"
	"net/smtp"
	"strings"
	"testing"
)

func TestEms(t *testing.T) {

	// SMTP 服务器信息
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// 发件人和收件人邮箱
	from := "liu15077731547@gmail.com"
	to := []string{"15077731547@163.com"}

	// 邮件内容
	subject := "Test Email from Go"
	body := "This is a test email sent from Go."

	// 连接 SMTP 服务器
	conn, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		panic(err)
	}

	// 进行 SMTP 身份验证 vlzfqsbomhcchwlo
	auth := smtp.PlainAuth("", from, "iydtrbsmucjnufsl", smtpHost)
	//if err = conn.Auth(auth); err != nil {
	//	panic(err)
	//}
	//
	//// 发送邮件
	//if err = conn.Mail(from); err != nil {
	//	panic(err)
	//}
	//
	//for _, recipient := range to {
	//	if err = conn.Rcpt(recipient); err != nil {
	//		panic(err)
	//	}
	//}
	//
	//// 发送邮件内容
	//w, err := conn.Data()
	//if err != nil {
	//	panic(err)
	//}

	msg := "From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	//if _, err = w.Write([]byte(msg)); err != nil {
	//	panic(err)
	//}
	//
	//if err = w.Close(); err != nil {
	//	panic(err)
	//}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		panic(err)
	}
	conn.Quit()

	fmt.Println("Email sent successfully!")

}
