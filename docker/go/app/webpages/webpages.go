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
	// テンプレートに出力する要素の構造体

	//要素Textsに構造体をおく(消しておk)
	// var tasks []dbctl.Task
	// dat.Texts = fmt.Sprint(tasks)

	dbctl.AddDB(r)
	// dat := data{}
	// dat.Texts = dbctl.CallDB()
	database := dbctl.CallDB()
	// fmt.Println(database)

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", database); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}
