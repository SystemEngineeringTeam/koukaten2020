package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Person はデータベースのテーブルから取得した値を扱うための構造体
type Person struct {
	ID     int
	Number string
	Name   string
}

func main() {
	db, err := sql.Open("mysql", "gopher:setsetset@tcp(mysql:3306)/sample")
	if err != nil {
		log.Println(err.Error())
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM persons")
	if err != nil {
		log.Println(err.Error())
		log.Println(err)
		os.Exit(2)
	}
	defer rows.Close()

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.HandleFunc("/", topPage)
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func topPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/index.html"))
	// テンプレートに出力する要素の構造体
	dat := struct {
		Texts string //型はなんでもOK(template.HTMLにするとHTMLタグも使える)、要素名はhtml側の出力枠に対応させる(1文字目は必ず大文字)
	}{
		Texts: "test", //入力フォームの下のテキスト
	}
	//要素Textsに構造体をおく(消しておk)
	hogePer := Person{114514, "hoge", "hogehoge"}
	dat.Texts = fmt.Sprint(hogePer)

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", dat); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}
