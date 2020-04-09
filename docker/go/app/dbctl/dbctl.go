package dbctl

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	// _ "github.com/go-sql-driver/mysql"
)

// Task はテンプレートに出力するための構造体です
type Task struct {
	ToDo    string
	Who     string
	Date    string
	Hours   string
	Minutes string
}

// AddDB はmain.goから呼び出してデータベースにデータを格納する関数です
func AddDB(r *http.Request) {
	r.ParseForm()
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	data := Task{r.FormValue("ToDo"), r.FormValue("Who"), r.FormValue("date"), r.FormValue("hours"), r.FormValue("minutes")}
	fmt.Println(data)

	// str := fmt.Sprintf("")

	ins, err := db.Prepare("insert into tasks(date,contents) values(?,?)")
	if err != nil {
		// log.Println(err, 3)
		// os.Exit(3)
		log.Fatal(err)
		return
	}
	defer ins.Close()

	ins.Exec(data.Date, data.ToDo)
	fmt.Println(ins)
}
