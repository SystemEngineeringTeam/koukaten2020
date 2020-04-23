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
	http.HandleFunc("/login", webpages.LoginPage)
	http.HandleFunc("/signup", webpages.SignUp)
	http.HandleFunc("/auth", webpages.AuthPage)
	http.HandleFunc("/presignup", webpages.PreSignUp)
	log.Println("Listening on :8080...")
	http.ListenAndServe(":80", nil)
}
