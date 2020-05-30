package main

import (
	"log"
	"net/http"

	"./auth"
	"./webpages"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.Handle("/htmlsrc/", http.StripPrefix("/htmlsrc", http.FileServer(http.Dir("html"))))

	http.HandleFunc("/", webpages.TopPage)

	http.HandleFunc("/login", webpages.LoginPage)
	http.HandleFunc("/signup", webpages.SignUp)
	http.HandleFunc("/signupComplete", webpages.SignUpComplete)

	http.HandleFunc("/auth", webpages.AuthPage)
	http.HandleFunc("/presignup", webpages.PreSignUp)

	http.HandleFunc("/book", webpages.BookDetails)
	http.HandleFunc("/add", webpages.BookAdd)
	http.HandleFunc("/search", webpages.SearchPage)

	http.HandleFunc("/user", webpages.UserPage)
	http.HandleFunc("/deletebook", webpages.BookDelete)

	http.HandleFunc("/logout", auth.Logout)

	// http.HandleFunc("/test", webpages.Test)
	log.Println("Listening on :8080...")
	http.ListenAndServe(":80", nil)
}
