package dbctl

import (
	"database/sql"
	"log"
	"os"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

// tableたちは使う時までコメントアウトしておく
// // Place はデータベースのテーブルから値を取得するための構造体
// type Place struct {
// 	PlaceID   int
// 	PlaceName string
// 	Name      string
// }

// // Book はデータベースのテーブルから値を取得するための構造体
// type Book struct {
// 	RFID         string
// 	BookName     string
// 	Isbn         string
// 	PlaceID      string
// 	BookDatetime string
// }

// // Person はデータベースのテーブルから値を取得するための構造体
// type Person struct {
// 	PersonID       int
// 	CardData       string
// 	PersonName     string
// 	PersonEmail    string
// 	Password       string
// 	PersonDatetime string
// }

// // BorrowedLog はデータベースのテーブルから値を取得するための構造体
// type BorrowedLog struct {
// 	BorrowedLogID int
// 	RfidTag       string
// 	PersonID      int
// }

// // PrePerson はデータベースのテーブルから値を取得するための構造体
// type PrePerson struct {
// 	PrePersonID       int
// 	PrePersonEmail    string
// 	PrePersonToken    string
// 	PrePersonDatetime string
// }

// PreRegister は仮登録データベースにメールアドレスとそのトークンと時刻を挿入する関数
func PreRegister(mail, token string) {
	//dbはめったに閉じる必要がないらしいがここでopenするのが適切かわからない
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	ins, err := db.Prepare("insert into pre_persons(pre_person_email,pre_person_token,pre_person_datetime) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ins.Close()

	ins.Exec(mail, token, time.Now().Format("2006-01-02 15:04:05"))
}

// CallAddress はメールアドレスを呼び出す関数
func CallAddress(token string) (address string) {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//pre_person_tokenとtokenが一致するpre_person_emailをrowsに格納する
	row, err := db.Query("select pre_person_email from pre_persons where pre_person_token = ?;", token)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&address)
	}

	return address
}

//PreUnRegister は本登録が完了した仮登録のレコードを削除する関数
func PreUnRegister(email string) {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//email が一致するレコードをdeleteする
	row, err := db.Query("delete from pre_persons where pre_person_email = '?';", email)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
}
