package auth

import (
	"log"
	"net/http"
)

// GetMail はCookieに保存されているメールアドレスを取得する関数
func GetMail(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("Mail")
	log.Println("mail", cookie.Value)

	if err != nil {
		log.Println(err)
	}

	return cookie.Value

}

// ChangeMailOfCookie はCookieに保存されているメールアドレスを書き換える関数
func ChangeMailOfCookie(w http.ResponseWriter, r *http.Request, mail string) {
	cookie, err := r.Cookie("Mail")
	if err != nil {
		log.Println(err)
	}
	cookie.Value = mail
	http.SetCookie(w, cookie)
}
