package webpages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	// "text/template"

	"../dbctl"
)

// TopPage はトップページを表示する関数です
// http.HandleFuncから呼び出して使います
func TopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	//フォームをパース
	r.ParseForm()

	//テンプレートをパース
	t := template.Must(template.ParseFiles("html/index.html"))

	dbctl.AddDB(r)
	database := dbctl.CallDB()

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", database); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}

}

// DeleteData はデータベースからタスクを削除する関数です
func DeleteData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dbctl.DeleteDB(r.FormValue("ID"))
	fmt.Println(r.FormValue("ID"))
	database := dbctl.CallDB()
	// fmt.Println(database)
	t := template.Must(template.ParseFiles("html/index.html"))

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", database); err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "localhost:8080", http.StatusMovedPermanently)
}

// SignUp は登録ページを表示するための関数です
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/SignUp.html"))
	// フォームを解析
	r.ParseForm()

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
	}

	// データベースにユーザーを追加する関数を呼び出す
	if err := dbctl.UserRegister(); err != nil {
		log.Println(err)
	}
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}

// AuthPage は認証ページを表示するための関数です
func AuthPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)

	// 表示するファイルを指定
	t := template.Must(template.ParseFiles("html/auth.html"))

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
	}
}
