package main

import (
	"log"
	"net/http"

	"./auth"
	"./webpages"
)

func main() {

	// fs := http.FileServer(http.Dir("html"))
	// http.Handle("/", fs)
	http.Handle("/htmlsrc/", http.StripPrefix("/htmlsrc", http.FileServer(http.Dir("html"))))

	http.HandleFunc("/", webpages.TopPage)

	http.HandleFunc("/login", webpages.LoginPage)
	http.HandleFunc("/signup", webpages.SignUp)
	http.HandleFunc("/signup/complete", webpages.SignUpComplete)

	http.HandleFunc("/auth", webpages.AuthPage)
	http.HandleFunc("/presignup", webpages.PreSignUp)

	http.HandleFunc("/book/detail", webpages.BookDetails)
	http.HandleFunc("/book/add", webpages.BookAdd)
	http.HandleFunc("/book/search", webpages.SearchPage)
	http.HandleFunc("/book/borrow", webpages.Borrow)
	http.HandleFunc("/book/delete", webpages.BookDelete)

	http.HandleFunc("/user", webpages.UserPage)
	http.HandleFunc("/user/setting", webpages.UserSetting)
	http.HandleFunc("/user/edit", webpages.UserEdit)

	http.HandleFunc("/logout", auth.Logout)

	http.HandleFunc("/test", webpages.Test)
	log.Println("Listening on :8080...")
	http.ListenAndServe(":80", nil)
}
