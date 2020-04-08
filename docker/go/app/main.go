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
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
	}
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, nil)
}
