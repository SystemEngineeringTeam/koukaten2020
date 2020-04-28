package dbctl

import (
	"database/sql"
	"log"
	"os"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

// tableたちは使う時までコメントアウトしておく これ使わないのでは？？？？
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

// Persons はUserRegisterの引数として用いる構造体
type Persons struct {
	CardData      string
	Name    string
	Email    string
	Password      string
}

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
	rows, err := db.Query("select pre_person_email from pre_persons where pre_person_token = ?;", token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&address)
	}

	return address
}

//PreUnRegister は本登録が完了した仮登録のレコードを削除する関数
func PreUnRegister(email string) {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	//email が一致するレコードをdeleteする
	row, err := db.Query("delete from pre_persons where pre_person_email = ?;", email)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()
}

//UserRegister はユーザを本登録する関数
func UserRegister(p Persons) error {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		log.Println("db",err)
		return err
	}
	defer db.Close()

	emailsRows,err:=db.Query("insert into emails (email,password) values (?,?);",p.Email,p.Password)
	if err != nil {
		log.Println("insert emailsRows",err)
		return err
	}
	
	// emailsテーブルからp.Emailとemailが一致するemail_idを取得する
	emailsRows,err=db.Query("select email_id from emails where email = ?",p.Email)
	if err != nil {
		log.Println("select emailsRows",err)
		return err
	}
	defer emailsRows.Close()

	emailsRows.Next()
	var emailID int
	err=emailsRows.Scan(&emailID)
	if err != nil{
		log.Println("Scan",err)
		return err
	}

	// 取得したemailIDを用いてpersonsにinsertする
	personsRows, err := db.Query("insert into persons (card_data,person_name,email_id,person_datetime) values (?,?,?,?);", p.CardData, p.Name, emailID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println(err)
		return err
	}
	defer personsRows.Close()

	return nil
}

// まだ修正してないよ
// // BookAdd はbooksに本を登録する関数
// func BookAdd(rfidTag, bookName, isbn, bookDatetime string, placeID int) error {
// 	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	defer db.Close()

// 	place := ""
// 	rows, err := db.Query("select * from places;")
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	fmt.Println(rows)

// 	for rows.Next() {
// 		rows.Scan(&place)
// 		log.Println("place:", place)
// 	}

// 	row, err := db.Query("insert into books (rfid_tag,book_name,isbn,place_id,book_datetime) values (?,?,?,?,?);", rfidTag, bookName, isbn, placeID, bookDatetime)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	defer row.Close()
// 	return err
// }
