package webpages

import (
	"fmt"
	"html/template"
	"net/http"

	// "text/template"

	"../dbctl"
)

type data struct {
	Texts []template.HTML
}

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

// Signup は登録ページを表示するための関数です
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	// フォームを解析
	r.ParseForm()

	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}
