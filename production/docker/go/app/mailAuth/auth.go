package mailauth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	token    string
}

// MailAuth はメール認証を行う関数です
func MailAuth(to, token string) {
	to = "tikuwamk2@gmail.com"

	m := mail{
		from:     "sysken.auth@gmail.com",
		username: "sysken.auth@gmail.com",
		password: "gqricdfchrthlnqd",
		to:       to,
		sub:      "メールアドレスの確認",
		msg:      "localhost:8080/auth?token=",
		token:    token,
	}

	if err := gmailSend(m); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (m mail) body() string {
	return "To: " + m.to + "\r\n" +
		"Subject: " + m.sub + "\r\b\r\n" +
		m.msg + m.token + "\r\n"
}

func gmailSend(m mail) error {
	smtpSvr := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com")
	if err := smtp.SendMail(smtpSvr, auth, m.from, []string{m.to}, []byte(m.body())); err != nil {
		return err
	}
	return nil
}

// GenerateToken はTokenを発行するための関数です
func GenerateToken(addr string) string {

	mail := "testaddr@gam.com"
	b := []byte(mail)

	// tokenの生成
	t := sha256.Sum256(b)

	token := hex.EncodeToString(t[:])

	fmt.Println("Generated : ", token)
	return token
}
