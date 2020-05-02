package webpages

import (
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"

	"../apictl"
	"../dbctl"
	mailauth "../mailauth"
)

// TopPage はトップページを表示する関数です
// http.HandleFuncから呼び出して使います
func TopPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/index.html"))

	//テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		log.Println(r.Form)
	}

}

//LoginPage はログインする時のページを表示する関数
func LoginPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)
	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/login.html"))

	dat := struct {
		Err template.HTML
	}{
		Err: "",
	}

	pwd := []byte(r.FormValue("Password"))
	hashedPassWord := sha256.Sum256(pwd)
	if ok, err := dbctl.Login(r.FormValue("User"), hex.EncodeToString(hashedPassWord[:])); ok {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		if r.FormValue("Password") != "" {
			dat.Err = `
			<br>
			<div class="alert alert-danger" role="alert">
				<p>` + template.HTML(err.Error()) + `</p>
			</div>
			`
			log.Println("err:", err)
		}
	}

	//テンプレートを描画
	if err := t.Execute(w, dat); err != nil {
		log.Println(err)
	}
}

// SignUp は登録ページを表示するための関数です
func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/signup.html"))
	// フォームを解析
	r.ParseForm()
	u := r.URL.Query()

	mail, err := dbctl.CallAddress(u["token"][0])
	if err != nil {
		log.Println(err)
	}

	// templateに渡す構造体を定義
	dat := struct {
		Mail string
	}{
		Mail: mail,
	}

	b := []byte(r.FormValue("Pass"))

	hashedPassWord := sha256.Sum256(b)

	User := dbctl.Persons{
		CardData: "hoge",
		Name:     r.FormValue("User"),
		Email:    dat.Mail,
		Password: hex.EncodeToString(hashedPassWord[:]),
	}

	// データベースにユーザーを追加する関数を呼び出す
	// 下の使用例を参照してUserRegister関数に適当な引数を入力してください
	if err := dbctl.UserRegister(User); err != nil {
		log.Println(err)
	}

	if r.Method == "POST" {
		log.Println(r.Form)
	}

	// テンプレートを描画
	if err := t.Execute(w, dat); err != nil {
		log.Println(err)
	}
}

// PreSignUp は仮登録ページを表示するための関数です
func PreSignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// テンプレートを指定
	t := template.Must(template.ParseFiles("html/presignup.html"))
	// t, err := template.New("PreSignUp").ParseFiles("html/presignup.html")
	// if err != nil {
	// 	log.Println(err)
	// }

	// フォームを解析
	if err := r.ParseForm(); err != nil {
		log.Println("form: ", err)
	}

	// 入力されたメールアドレスを取得
	mail := r.FormValue("Mail")
	log.Println("Mail: ", mail)

	dat := struct {
		Msg template.HTML
	}{
		Msg: "",
	}

	// メアドが入力されていればメールを送信する
	if mail != "" {
		// 認証メールを送信する関数にメールアドレスを渡す
		mailauth.MailAuth(mail)
		dat.Msg = `
			<br>
			<div class="alert alert-success" role="alert">
				<p>仮登録メールを送信しました</p>
			</div>
		`
	}

	// テンプレートを描画
	if err := t.Execute(w, dat); err != nil {
		log.Println(err)
	}
}

// AuthPage は認証ページを表示するための関数です
func AuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/auth.html"))

	// log.Println(r.URL)
	u := r.URL.Query()
	log.Println(u["token"])

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
}

// BookDetails は本の詳細ページを表示するための関数です
func BookDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/bookPage.html"))

	// log.Println(r.URL)
	u := r.URL.Query()
	log.Println(u["isbn"][0])

	detail := apictl.BookDetail(u["isbn"][0])

	// テンプレートを描画
	if err := t.Execute(w, detail); err != nil {
		log.Println(err)
	}

}

//SignUpComplete はユーザー登録完了ページの関数です
func SignUpComplete(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/signupComplete.html"))

	// log.Println(r.URL)
	u := r.URL.Query()
	log.Println(u["token"])

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

//SearchPage は本の一覧ページ
func SearchPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/bookSerch.html"))

	r.ParseForm()

	searchedBooks := apictl.SearchBooks(r.FormValue("searchForm"))

	// テンプレートを描画
	if err := t.Execute(w, searchedBooks); err != nil {
		log.Println(err)
	}

}

//Test は新しく作った関数をテストするところ 関数の使い方も兼ねている
func Test(w http.ResponseWriter, r *http.Request) {

	// log.Println(dbctl.Login("e19070ee@aitech.ac.jp", "4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d"))
}
