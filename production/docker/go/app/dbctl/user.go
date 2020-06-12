package dbctl

import (
	"fmt"
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
	u := User{}
	var emailID int

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
	fmt.Println("hogehoge", u)

	return u, nil
}
