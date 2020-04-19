package mailauth

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
func MailAuth(to string) {

	m := mail{
		from:     "sysekn.auth@gmail.com",
		username: "sysken.auth@gmail.com",
		password: "gqricdfchrthlnqd",
		to:       to,
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

// GenerateToken はTokenを発行するための関数です
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	signBytes, err := ioutil.ReadFile("./app_rsa")
	if err != nil {
		log.Fatal(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	// tokenの生成
	token := jwt.New(jwt.SigningMethodRS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "test"
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}
