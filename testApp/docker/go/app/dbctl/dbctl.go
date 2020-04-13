package dbctl

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	// _ "github.com/go-sql-driver/mysql"
)

// Person はデータベースのテーブルから取得した値を扱うための構造体
type Person struct {
	ID     int
	Number string
	Name   string
}

// Task はテンプレートに出力するための構造体です
type Task struct {
	ToDo    string
	Who     string
	Date    string
	Hours   string
	Minutes string
}

// DoPut はデータベースから値を取得するための構造体
type DoPut struct {
	ID       string
	DateTime string
	PersonID string
	Who      Person
	Contents string
}

// AddDB はmain.goから呼び出してデータベースにデータを格納する関数です
func AddDB(r *http.Request) {
	r.ParseForm()
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	data := Task{r.FormValue("ToDo"), r.FormValue("Who"), r.FormValue("date"), r.FormValue("hours"), r.FormValue("minutes")}
	// fmt.Println(data)

	// str := fmt.Sprintf("")

	ins, err := db.Prepare("insert into tasks(datetime,person_id,contents) values(?,?,?)")
	if err != nil {
		// log.Println(err, 3)
		// os.Exit(3)
		log.Fatal(err)
		return
	}
	defer ins.Close()

	ins.Exec(data.Date+" "+data.Hours+":"+data.Minutes, data.Who, data.ToDo)
	fmt.Println(ins)
}

// CallDB はデータベースから値を取得します
func CallDB() []DoPut {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("select * from tasks;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// var database []DoPut
	database := make([]DoPut, 0)

	for rows.Next() {
		tpl := DoPut{"", "", "", Person{0, "", ""}, ""}
		err = rows.Scan(&tpl.ID, &tpl.DateTime, &tpl.PersonID, &tpl.Contents)
		if err != nil {
			log.Println(err)
		}

		tpl.Who = getPerson(tpl.PersonID, db)
		fmt.Println("DB", tpl.ID, tpl.DateTime, getPerson(tpl.PersonID, db).string(), tpl.Contents)
		// fmt.Println(tpl.DateTime)
		database = append(database, tpl)
	}
	return database
}

func getPerson(p string, db *sql.DB) Person {
	human := Person{}
	rows, err := db.Query("select * from persons where id = ?;", p)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&human.ID, &human.Number, &human.Name)
		if err != nil {
			fmt.Println(err)
		}
	}

	return human
}

func (p Person) string() string {
	return fmt.Sprintf("%s %s", p.Number, p.Name)
}

// DeleteDB はタスクのIDを指定してタスクを削除する関数です
func DeleteDB(id string) {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		// log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("delete from tasks where id = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)
}
