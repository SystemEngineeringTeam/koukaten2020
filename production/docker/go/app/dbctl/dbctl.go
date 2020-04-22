package dbctl

import (
	"database/sql"
	"log"
	"os"
	// _ "github.com/go-sql-driver/mysql"
)

// Place はデータベースのテーブルから値を取得するための構造体
type Place struct {
	PlaceID   int
	PlaceName string
	Name      string
}

// Book はデータベースのテーブルから値を取得するための構造体
type Book struct {
	RFID         string
	BookName     string
	Isbn         string
	PlaceID      string
	BookDatetime string
}

// Person はデータベースのテーブルから値を取得するための構造体
type Person struct {
	PersonID       int
	CardData       string
	PersonName     string
	PersonEmail    string
	Password       string
	PersonDatetime string
}

// BorrowedLog はデータベースのテーブルから値を取得するための構造体
type BorrowedLog struct {
	BorrowedLogID int
	RfidTag       string
	PersonID      int
}

// PrePerson はデータベースのテーブルから値を取得するための構造体
type PrePerson struct {
	PrePersonID       int
	PrePersonEmail    string
	PrePersonToken    string
	PrePersonDatetime string
}

// PreRegister は仮登録データベースにメールアドレスとそのトークンを登録する関数
func PreRegister(mail, token string) {
	mail = "hello@gmail.com"
	token = "1d945e4947da1a05bf393b67b2e0a1fe2be36965cd4f44da5069a1df505e0092"
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	ins, err := db.Prepare("insert into pre_persons(pre_person_email,pre_person_token,pre_person_datetime) values(?,?,?)")
	if err != nil {
		// log.Println(err, 3)
		// os.Exit(3)
		log.Fatal(err)
		return
	}
	defer ins.Close()

	ins.Exec(mail, token, "20200422194000")
}
