package main

import (
	"log"
	"net/http"

	"./webpages"

	_ "github.com/go-sql-driver/mysql"
)

// Person はデータベースのテーブルから取得した値を扱うための構造体
type Person struct {
	ID     int
	Number string
	Name   string
}

func main() {

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.HandleFunc("/", webpages.TopPage)
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
