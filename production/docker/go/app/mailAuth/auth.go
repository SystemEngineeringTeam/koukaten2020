package main

import (
	"log"
	"net/smtp"
	"os"
)

type mail struct {
	from     string
	username string
	password string
	to       string
	sub      string
	msg      string
}

// MailAuth はメール認証を行う関数です
func main() {

	m := mail{
		from:     "sysekn.auth@gmail.com",
		username: "sysken.auth@gmail.com",
		password: "gqricdfchrthlnqd",
		to:       "tikuwamk2@gmail.com",
		sub:      "メールアドレスの確認",
		msg:      "",
	}

	if err := gmailSend(m); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (m mail) body() string {
	return "To: " + m.to + "\r\n" +
		"Subject: " + m.sub + "\r\b\r\n" +
		m.msg + "\r\n"
}

func gmailSend(m mail) error {
	smtpSvr := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com")
	if err := smtp.SendMail(smtpSvr, auth, m.from, []string{m.to}, []byte(m.body())); err != nil {
		return err
	}
	return nil
}
