package mailauth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"

	"../dbctl"
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
func MailAuth(to string) {

	m := mail{
		from:     "sysken.auth@gmail.com",
		username: "sysken.auth@gmail.com",
		password: "gqricdfchrthlnqd",
		to:       to,
		sub:      "メールアドレスの確認",
		msg:      "localhost:8080/auth?token=",
		token:    generateToken(to),
	}

	if err := gmailSend(m); err != nil {
		log.Println(err)

	}

	dbctl.PreRegister(to, m.token)
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

func generateToken(addr string) string {
	b := []byte(addr)

	// tokenの生成
	t := sha256.Sum256(b)

	token := hex.EncodeToString(t[:])

	fmt.Println("Generated :", token)
	return token
}
