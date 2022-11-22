package utils

import (
	"github.com/ChocolateAceCream/blog/global"
	"gopkg.in/gomail.v2"
)

func SendMail(email, subject, body string) (err error) {
	msg := gomail.NewMessage()
	config := global.CONFIG.Email
	msg.SetHeader("From", config.Username)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	n := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	// Send the email
	return n.DialAndSend(msg)
}
