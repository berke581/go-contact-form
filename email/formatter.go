package email

import (
	"fmt"
)

func formatMessage(to string, name string, title string, email string, body string) []byte {
	var from string = ""
	if len(name) > 0 {
		from += fmt.Sprintf("%s ", name)
	}
	from += fmt.Sprintf("<%s>", email)

	return []byte(fmt.Sprintf("From: %s\r\n", from) +
		fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n", title) +
		"\r\n" +
		fmt.Sprintf("%s\r\n", body))
}
