package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/gorilla/sessions"

	"../dbctl"
)

// ここからメール認証

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
		msg:      "localhost:8080/signup?token=",
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

func mailToHashedMail(mail string) string {
	b := sha256.Sum256([]byte(mail))
	hashedMail := hex.EncodeToString(b[:])
	return hashedMail
}

// ここからセッション管理

var store = sessions.NewCookieStore([]byte("setsetset"))

// CreateNewSession は新しいセッションを作成する関数です
func CreateNewSession(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mail := r.FormValue("Mail")
	hashedMail := mailToHashedMail(mail)
	session, _ := store.Get(r, hashedMail)
	session.Values["login"] = true
	session.Save(r, w)

}

// IsLogin はログイン状態を判別する関数です
func IsLogin(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	mail := r.FormValue("Mail")
	hashedMail := mailToHashedMail(mail)
	session, _ := store.Get(r, hashedMail)

	if session.Values["login"] == true {
		return true
	}
	return false
}