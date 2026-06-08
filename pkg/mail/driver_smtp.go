package mail

import (
	"fmt"
	"gohub/pkg/logger"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("Send email", "email", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)

	if err != nil {
		logger.ErrorString("Send email", "error", err.Error())
		return false
	}

	logger.DebugString("Send email", "success", "")

	return true
}
