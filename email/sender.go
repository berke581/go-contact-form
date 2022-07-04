package email

import (
	"fmt"
	"log"
	"net/smtp"
)

type emailSender struct {
	from     string
	password string
}

func NewEmailSender(from string, password string) *emailSender {
	return &emailSender{
		from,
		password,
	}
}

var smtpHost = "smtp.gmail.com"
var smtpPort = "587"

func (emailSender *emailSender) SendEmail(to string, title string, msg string) {
	from := emailSender.from
	pw := emailSender.password

	toArray := []string{to}
	message := []byte(
		fmt.Sprintf("Subject: %s\r\n\r\n", title) +
			fmt.Sprintf("%s\r\n", msg))
	auth := smtp.PlainAuth("", from, pw, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toArray, message)
	if err != nil {
		log.Fatalln("Error sending mail: ", err)
	}

	log.Println("Email sent succesfully.")
}
