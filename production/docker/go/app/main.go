package main

import (
	"log"
	"net/http"

	mailauth "./mailAuth"
	"./webpages"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.HandleFunc("/", webpages.TopPage)
	http.HandleFunc("/delete", webpages.DeleteData)
	http.HandleFunc("/signup", webpages.SignUp)
	http.HandleFunc("/auth", auth)
	log.Println("Listening on :8080...")
	http.ListenAndServe(":80", nil)
}

func auth(w http.ResponseWriter, r *http.Request) {
	mailauth.GenerateToken(w, r)
}
