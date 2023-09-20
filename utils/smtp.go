package utils

import (
	"fmt"
	"net/smtp"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
)

var (
	auth smtp.Auth
	conf config.MailConf
)

func InitMail() {
	conf = config.GetMailConfig()
	auth = smtp.PlainAuth("", conf.FromAddress, conf.Password, conf.Host)
}

func EasySendMail(sub, body, to string) error {
	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+body,
		conf.FromAddress, to, sub,
	)
	return smtp.SendMail(conf.Server, auth, conf.FromAddress, []string{to}, []byte(msg))
}
