package webpages

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../apictl"
	"../auth"
	"../dbctl"
)

// TopPage はトップページを表示する関数です
// http.HandleFuncから呼び出して使います
func TopPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	if auth.IsLogin(w, r) == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	books, err := dbctl.BookStatus()
	if err != nil {
		log.Println(err)
		return
	}

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/index.html"))

	//テンプレートを描画
	if err := t.Execute(w, books); err != nil {
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
		auth.CreateNewSession(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
		Err  template.HTML
	}{
		Mail: mail,
		Err:  ``,
	}

	b := []byte(r.FormValue("Pass"))

	hashedPassWord := sha256.Sum256(b)
	c := sha256.Sum256([]byte(r.FormValue("User")))
	User := dbctl.Persons{
		CardData: hex.EncodeToString(c[:]),
		Name:     r.FormValue("User"),
		Email:    dat.Mail,
		Password: hex.EncodeToString(hashedPassWord[:]),
	}
	// データベースにユーザーを追加する関数を呼び出す
	// 下の使用例を参照してUserRegister関数に適当な引数を入力してください
	if User.Name != "" && User.CardData != "" && User.Email != "" && User.Password != "" {
		if err := dbctl.UserRegister(User); err != nil {
			log.Println(err)
			dat.Err = `
			<br>
			<div class="alert alert-danger" role="alert">
				<p>エラーが発生しました</p>
			</div>
		`
		}
		http.Redirect(w, r, "/signup/complete", http.StatusSeeOther)
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
		auth.MailAuth(mail)
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

	http.Redirect(w, r, "/signup", http.StatusSeeOther)
}

// BookDetails は本の詳細ページを表示するための関数です
func BookDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/bookPage.html"))

	// log.Println(r.URL)
	u := r.URL.Query()
	log.Println(u["id"][0])

	detail, err := dbctl.BookDetail(u["id"][0])
	if err != nil {
		log.Println(err)
	}

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
	t := template.Must(template.ParseFiles("html/bookSearch.html"))

	r.ParseForm()

	searchedBooks := apictl.SearchBooks(r.FormValue("searchForm"))

	// テンプレートを描画
	if err := t.Execute(w, searchedBooks); err != nil {
		log.Println(err)
	}

}

// BookAdd は本を追加するためのページ
// 検索ページから飛んでくる
func BookAdd(w http.ResponseWriter, r *http.Request) {

	log.Println("URL:", r.URL)

	u := r.URL.Query()
	log.Println("u:", u)
	log.Println("len(u[\"id\"]):", len(u["id"]))

	if len(u["id"]) > 0 {
		fmt.Println("hogehoge")
		b := apictl.BookRegister(u["id"][0])

		fmt.Println(b)

		if b.RFID != "error" {
			err := dbctl.BookAdd(b)
			if err != nil {
				log.Println("hoge")
				log.Println(err)
			}
		}

	} else {
		log.Println("id is undefined.")

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// UserPage はユーザー情報を閲覧するページの関数
func UserPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/userPage.html"))

	// ユーザーの情報を取ってくる関数
	mail := auth.GetMail(w, r)

	// ユーザーの情報を表示するための構造体
	u, err := dbctl.CallUserFromMail(mail)
	if err != nil {
		log.Println(err)
		return
	}
	// テンプレートを描画
	if err := t.Execute(w, u); err != nil {
		log.Println(err)
		return
	}

}

// BookDelete は本の削除ボタンに使う関数です
func BookDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL)

	r.ParseForm()
	if err := dbctl.DeleteBook(r.FormValue("DeleteID")); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// FavHandle は/favicon.icoに対する処理を記述した関数です
func FavHandle(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w,r,"relative/path/to/favicon.ico")
}

// Borrow は本を借りる処理を行う関数です
func Borrow(w http.ResponseWriter, r *http.Request) {
	// 引数はRFID、借りた人の学生証の値
	r.ParseForm()
	err := dbctl.BorrowBook(r.FormValue("RFID"), "hoge")
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// UserSetting はユーザーの情報を変更するためにフォーム入力させるページ
func UserSetting(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/userEdit.html"))

	// ユーザーの情報を取ってくる関数
	mail := auth.GetMail(w, r)

	// ユーザーの情報を表示するための構造体
	u, err := dbctl.CallUserFromMail(mail)
	if err != nil {
		log.Println(err)
		return
	}

	if err := t.Execute(w, u); err != nil {
		log.Println(err)
	}

}

// UserEdit はユーザ情報の編集をするページ
func UserEdit(w http.ResponseWriter, r *http.Request) {
	// フォームの解析
	r.ParseForm()

	// ユーザーの情報を取ってくる関数
	mail := auth.GetMail(w, r)

	// ユーザーの情報を表示するための構造体
	u, err := dbctl.CallUserFromMail(mail)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("user", u)
	fmt.Println(r.FormValue("formUser"), r.FormValue("formMail"), r.FormValue("formPass"))

	// 各情報を変更する関数
	if ok := r.FormValue("User"); ok != "" {
		err := u.ChangeUserName(ok)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if ok := r.FormValue("Mail"); ok != "" {
		err := u.ChangeEmail(ok)
		if err != nil {
			log.Println(err)
			return
		}
		auth.ChangeMailOfCookie(w, r, ok)
	}
	if ok := r.FormValue("Pass"); ok != "" {
		err := u.ChangePassword(ok)
		if err != nil {
			log.Println(err)
			return
		}
	}

	http.Redirect(w, r, "/user", http.StatusSeeOther)

}

//Test は新しく作った関数をテストするところ 関数の使い方も兼ねている
func Test(w http.ResponseWriter, r *http.Request) {
	// 引数はRFID、借りた人の学生証の値
	log.Println(dbctl.BorrowBook("hoge", "hoge"))
}
