package utils

import (
	"fmt"
	"log"
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

func EasySendMail(sub, body, to string) {
	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+body,
		conf.FromAddress, to, sub,
	)
	err := smtp.SendMail(conf.Server, auth, conf.FromAddress, []string{to}, []byte(msg))
	if err != nil {
		log.Println("mailer:", err)
	}
}
