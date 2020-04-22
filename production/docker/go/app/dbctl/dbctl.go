package dbctl

import (
	"database/sql"
	"log"
	"os"
	"time"
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

<<<<<<< HEAD
<<<<<<< HEAD
// Account はアカウントを管理する構造体です
type Account struct {
	Name     string
	Mail     string
	Password string
	Token    string
}

// AddDB はmain.goから呼び出してデータベースにデータを格納する関数です
func AddDB(r *http.Request) {
	r.ParseForm()
=======
// PreRegister は仮登録データベースにメールアドレスとそのトークンを登録する関数
func PreRegister(mail, token string) {
	mail = "hello@gmail.com"
	token = "1d945e4947da1a05bf393b67b2e0a1fe2be36965cd4f44da5069a1df505e0092"
>>>>>>> dev_db
=======
// PreRegister は仮登録データベースにメールアドレスとそのトークンと時刻を登録する関数
func PreRegister(mail, token string) {
>>>>>>> dev_db
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

	ins.Exec(mail, token, time.Now().Format("2006-01-02 15:04:05"))
}

// UserRegister はユーザー登録の際のDB処理を行う関数です
func UserRegister() error {
	return nil
}
