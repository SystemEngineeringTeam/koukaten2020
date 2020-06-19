package dbctl

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"runtime"
)

// User はユーザー情報を扱うための構造体です
type User struct {
	CardData string
	Name     string
	Email    string
	Datetime string
}

// CallUserFromMail はメールアドレスからユーザー情報を呼び出す関数です
func CallUserFromMail(mail string) (User, error) {
	// User構造体の定義
	u := User{}
	// DBのIDを定義
	emailID := callEmailID(mail)

	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	rows, err := db.Query("select email_id from emails where email = ?", mail)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return u, err
	}
	defer rows.Close()

	rows.Next()
	rows.Scan(&emailID)

	// person_id | card_data | person_name | email_id | person_datetime
	users, err := db.Query("select card_data,person_name,person_datetime from persons where email_id = ?", emailID)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return u, err
	}
	defer users.Close()
	users.Next()
	users.Scan(&u.CardData, &u.Name, &u.Datetime)
	u.Email = mail

	return u, nil
}

// callEmailID はEmailIDを呼び出す関数
func callEmailID(mail string) int {
	var emailID int

	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	rows, err := db.Query("select email_id from emails where email = ?", mail)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1
	}
	defer rows.Close()

	rows.Next()
	rows.Scan(&emailID)
	return emailID
}

// ChangeUserName はユーザー名を編集するメソッド
func (u User) ChangeUserName(name string) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	rows, err := db.Query("update persons set person_name = ? where email_id = ?", name, callEmailID(u.Email))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	u.Name = name
	return nil
}

// ChangeEmail はユーザー名を編集するメソッド
func (u User) ChangeEmail(mail string) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	rows, err := db.Query("update emails set email = ? where email_id = ?", mail, callEmailID(u.Email))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	u.Email = mail
	return nil
}

// ChangePassword はユーザー名を編集するメソッド
func (u User) ChangePassword(pass string) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	p := sha256.Sum256([]byte(pass))
	hashedPassword := hex.EncodeToString(p[:])
	rows, err := db.Query("update emails set password=? where email_id = ?", hashedPassword, callEmailID(u.Email))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	return nil
}
