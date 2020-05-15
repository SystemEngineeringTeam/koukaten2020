package dbctl

import (
	"database/sql"
	"errors"
	"log"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//エラーの内容:err 関数の名前:f.Name() ファイルのパス:file runtimeが呼ばれた行数:line
const errFormat = "\n%v\nfunction:%v file:%v line:%v\n"

// // Place はデータベースのテーブルから値を取得するための構造体
// type Place struct {
// 	PlaceID   int
// 	PlaceName string
// 	Name      string
// }

// Book は本の登録、詳細な情報の表示に使用する構造体
type Book struct {
	RFID          string
	PlaceID       int
	Status        string
	BookName      string
	APIID         string
	Author        string
	Publisher     string
	PublishedDate string
	Description   string
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

var db *sql.DB

// packageがimportされたときに呼び出される関数
func init() {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	var err error

	db, err = sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/book_management_db")
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)

		//データベースを開けないと動作が継続できないためpanicを発生させる
		panic("Can't Open database.")
	}
}

// PreRegister は仮登録データベースにメールアドレスとそのトークンと時刻を挿入する関数
func PreRegister(mail, token string) error {
	// この関数の名前とファイルのパスとruntimeが呼ばれた行数を取得する
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	rows, err := db.Query("insert into pre_persons(pre_person_email,pre_person_token,pre_person_datetime) values(?,?,?)", mail, token, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	return err
}

// CallAddress はメールアドレスを呼び出す関数
func CallAddress(token string) (address string, err error) {
	// エラーの時のaddressに格納される文字列
	errStr := "Can't call address"

	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	//pre_person_tokenとtokenが一致するpre_person_emailをrowsに格納する
	rows, err := db.Query("select pre_person_email from pre_persons where pre_person_token = ?;", token)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return errStr, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&address)
	}

	return address, err
}

//PreUnRegister は本登録が完了した仮登録のレコードを削除する関数
func PreUnRegister(email string) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	//email が一致するレコードをdeleteする
	row, err := db.Query("delete from pre_persons where pre_person_email = ?;", email)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer row.Close()

	return err
}

//UserRegister はユーザを本登録する関数
func UserRegister(p Persons) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	emailsRows, err := db.Query("insert into emails (email,password) values (?,?);", p.Email, p.Password)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	// emailsテーブルからp.Emailとemailが一致するemail_idを取得する
	emailsRows, err = db.Query("select email_id from emails where email = ?", p.Email)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer emailsRows.Close()

	emailsRows.Next()
	var emailID int
	err = emailsRows.Scan(&emailID)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	// 取得したemailIDを用いてpersonsにinsertする
	personsRows, err := db.Query("insert into persons (card_data,person_name,email_id,person_datetime) values (?,?,?,?);", p.CardData, p.Name, emailID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer personsRows.Close()

	return err
}

// BookAdd はbook_info,book_statuesに本を登録する関数
func BookAdd(b Book) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	bookInfoRows, err := db.Query("insert into book_info (book_name,api_id,author,publisher,published_date,description) values (?,?,?,?,?,?);", b.BookName, b.APIID, b.Author, b.Publisher, b.PublishedDate, b.Description)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	bookInfoRows, err = db.Query("select book_info_id from book_info where book_name = ?;", b.BookName)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer bookInfoRows.Close()

	bookInfoRows.Next()
	var bookInfoID int
	err = bookInfoRows.Scan(&bookInfoID)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	bookStatusesRows, err := db.Query("insert into book_statuses (rfid_tag,book_info_id,place_id,book_datetime) values (?,?,?,?);", b.RFID, bookInfoID, b.PlaceID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer bookStatusesRows.Close()

	return err
}

// BookStatus はplaceIDが示す場所に存在する本の情報を返す関数
func BookStatus(placeID int) ([]Book, error) {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	errStr := "This place is error."
	// placeIDが0,-1は貸出中持出中であるのでエラー
	if placeID == 0 || placeID == -1 {
		log.Printf(errFormat, errStr, f.Name(), file, line)
		return nil, errors.New(errStr)
	}

	books := make([]Book, 0, 10)
	var bookBuf Book
	// その場所(placeIDに該当する)にある本がいくつかを数えるためにbooksとは別のスライスにしている
	bookInfoIDs := make([]int, 0, 10)
	var infoIDBuf int

	// placeIDが一致する本のレコードをbook_statusesからselectする
	booksStatusRows, err := db.Query("select rfid_tag,book_info_id,place_id from book_statuses where place_id = ?;", placeID)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer booksStatusRows.Close()

	for booksStatusRows.Next() {
		err := booksStatusRows.Scan(&bookBuf.RFID, &infoIDBuf, &bookBuf.PlaceID)
		if err != nil {
			log.Printf(errFormat, err, f.Name(), file, line)
			return nil, err
		}

		// 一時的に格納したレコードを追加する
		books = append(books, bookBuf)
		bookInfoIDs = append(bookInfoIDs, infoIDBuf)
	}

	// booksとbookinfoIDは一対一に対応しているため、forのindexが示すbooksの要素とIDを引数としてselectしたレコードは同じ本の情報となる
	for i, ID := range bookInfoIDs {
		booksInfoRows, err := db.Query("select book_name,api_id,author,publisher,published_date,description from book_info where book_info_id = ?;", ID)
		if err != nil {
			log.Printf(errFormat, err, f.Name(), file, line)
			return nil, err
		}
		defer booksInfoRows.Close()

		booksInfoRows.Next()
		err = booksInfoRows.Scan(&books[i].BookName, &books[i].APIID, &books[i].Author, &books[i].Publisher, &books[i].PublishedDate, &books[i].Description)
		if err != nil {
			log.Printf(errFormat, err, f.Name(), file, line)
			return nil, err
		}
	}

	return books, err
}

// Login はメアドとパスワードでログインする関数です
func Login(mail, pass string) (bool, error) {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	errStr := "メールアドレスまたはパスワードが間違っています"

	rows, err := db.Query("select email_id from emails where email = ? and password = ?;", mail, pass)
	if err != nil {
		log.Println(errFormat, err, f.Name(), file, line)
		return false, errors.New(errStr)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, errors.New(errStr)
}
