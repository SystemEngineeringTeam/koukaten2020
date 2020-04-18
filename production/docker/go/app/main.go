package main

import (
	"log"
	"net/http"

	"./webpages"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.HandleFunc("/", webpages.TopPage)
	http.HandleFunc("/delete", webpages.DeleteData)
	http.HandleFunc("/signup", webpages.Signup)
	log.Println("Listening on :8080...")
	http.ListenAndServe(":80", nil)
}
