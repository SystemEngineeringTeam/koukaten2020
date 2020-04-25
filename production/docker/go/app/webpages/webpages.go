package webpages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	mailauth "../mailAuth"
)

// TopPage はトップページを表示する関数です
// http.HandleFuncから呼び出して使います
func TopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/header.html"))

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", nil); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}

}

//LoginPage はログインする時のページを表示する関数
func LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/login.html"))

	//テンプレートを描画
	if err := t.Execute(w, "nil"); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}

// SignUp は登録ページを表示するための関数です
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/signup.html"))
	// フォームを解析
	r.ParseForm()

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
	}

	// データベースにユーザーを追加する関数を呼び出す
	// if err := dbctl.UserRegister(); err != nil {
	// 	log.Println(err)
	// }
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}

// PreSignUp は仮登録ページを表示するための関数です
func PreSignUp(w http.ResponseWriter, r *http.Request) {
	// テンプレートを指定
	t := template.Must(template.ParseFiles("html/presignup.html"))

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}

	// フォームを解析
	r.ParseForm()

	// 入力されたメールアドレスを取得
	mail := r.FormValue("Mail")
	fmt.Println("Mail: ", mail)

	if mail != "" {

		// 認証メールを送信する関数にメールアドレスを渡す
		mailauth.MailAuth(mail)
		http.Redirect(w, r, "/presignup", http.StatusMovedPermanently)
	}

}

// AuthPage は認証ページを表示するための関数です
func AuthPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/auth.html"))

	// fmt.Println(r.URL)
	u := r.URL.Query()
	fmt.Println(u["token"])

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
}
