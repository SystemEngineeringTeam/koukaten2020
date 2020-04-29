package dbctl

import (
	"database/sql"
	"log"
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

// Book は本の登録、情報の表示に使用する構造体
type Book struct {
	RFID          string
	Status        string
	PlaceID       string
	BookName      string
	Authors       string
	Publisher     string
	PublishedDate string
	ISBN          string
}

// Persons はUserRegisterの引数として用いる構造体
type Persons struct {
	CardData string
	Name     string
	Email    string
	Password string
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
func PreRegister(mail, token string) error {
	//dbはめったに閉じる必要がないらしいがここでopenするのが適切かわからない
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Println("PreRegister:Open", err)
		return err
	}
	defer db.Close()

	rows, err := db.Query("insert into pre_persons(pre_person_email,pre_person_token,pre_person_datetime) values(?,?,?)", mail, token, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("PreRegister:insert", err)
		return err
	}
	defer rows.Close()

	return nil
}

// CallAddress はメールアドレスを呼び出す関数
func CallAddress(token string) (address string, err error) {
	errStr := "This is error!"

	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Println("CallAddress:Open", err)
		return errStr, err
	}
	defer db.Close()

	//pre_person_tokenとtokenが一致するpre_person_emailをrowsに格納する
	rows, err := db.Query("select pre_person_email from pre_persons where pre_person_token = ?;", token)
	if err != nil {
		log.Println("CallAddress:select:", err)
		return errStr, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&address)
	}

	return address, nil
}

//PreUnRegister は本登録が完了した仮登録のレコードを削除する関数
func PreUnRegister(email string) error {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Println("PreUnRegister:Open", err)
		return err
	}
	defer db.Close()

	//email が一致するレコードをdeleteする
	row, err := db.Query("delete from pre_persons where pre_person_email = ?;", email)
	if err != nil {
		log.Println("PreUnRegister:delete", err)
		return err
	}
	defer row.Close()

	return err
}

//UserRegister はユーザを本登録する関数
func UserRegister(p Persons) error {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Println("UserRegister:Open", err)
		return err
	}
	defer db.Close()

	emailsRows, err := db.Query("insert into emails (email,password) values (?,?);", p.Email, p.Password)
	if err != nil {
		log.Println("UserRegister:insert emails", err)
		return err
	}

	// emailsテーブルからp.Emailとemailが一致するemail_idを取得する
	emailsRows, err = db.Query("select email_id from emails where email = ?", p.Email)
	if err != nil {
		log.Println("UserRegister:select", err)
		return err
	}
	defer emailsRows.Close()

	emailsRows.Next()
	var emailID int
	err = emailsRows.Scan(&emailID)
	if err != nil {
		log.Println("UserRegister:Scan", err)
		return err
	}

	// 取得したemailIDを用いてpersonsにinsertする
	personsRows, err := db.Query("insert into persons (card_data,person_name,email_id,person_datetime) values (?,?,?,?);", p.CardData, p.Name, emailID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("UserRegister:insert persons", err)
		return err
	}
	defer personsRows.Close()

	return nil
}

// BookAdd はbook_info,book_statusに本を登録する関数
func BookAdd(b Book) error {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Println("BookAdd:Open", err)
		return err
	}
	defer db.Close()

	bookInfoRows, err := db.Query("insert into book_info (book_name,authors,publisher,published_date,description,isbn) values (?,?,?,?,?,?);", b.BookName, b.Authors, b.Publisher, b.PublishedDate, b.ISBN)
	if err != nil {
		log.Println("BookAdd:insert book_info")
		return err
	}

	bookInfoRows, err = db.Query("select book_info_id from book_info where book_name = ?;", b.BookName)
	if err != nil {
		log.Println("BookAdd:book_info")
		return err
	}
	defer bookInfoRows.Close()

	bookInfoRows.Next()
	var bookInfoID int
	err = bookInfoRows.Scan(&bookInfoID)
	if err != nil {
		log.Println("BookAdd:Scan", err)
		return err
	}

	bookStatusesRows, err := db.Query("insert into book_statuses (rfid_tag,book_info_id,status,place_id,book_datetime) values (?,?,?,?,?);", b.RFID, bookInfoID, b.Status, b.PlaceID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("BookAdd:insert", err)
		return err
	}
	defer bookStatusesRows.Close()
	return err
}
