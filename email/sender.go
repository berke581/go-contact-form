package email

import (
	"log"
	"net/smtp"
	"os"
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

func (emailSender *emailSender) SendEmail(to string, name string, title string, email string, body string) {
	from := emailSender.from
	password := emailSender.password

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	toList := []string{to}
	message := formatMessage(to, name, title, email, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toList, message)
	if err != nil {
		log.Fatalln("Error sending mail: ", err)
	}

	log.Println("Email sent succesfully.")
}
