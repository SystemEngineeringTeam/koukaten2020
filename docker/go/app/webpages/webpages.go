package webpages

import (
	"fmt"
	"net/http"
	"text/template"

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
	// テンプレートに出力する要素の構造体
	dat := struct {
		Texts string //型はなんでもOK(template.HTMLにするとHTMLタグも使える)、要素名はhtml側の出力枠に対応させる(1文字目は必ず大文字)
	}{
		Texts: "test", //入力フォームの下のテキスト
	}
	//要素Textsに構造体をおく(消しておk)
	// var tasks []dbctl.Task
	// dat.Texts = fmt.Sprint(tasks)

	dbctl.AddDB(r)

	//テンプレートを描画
	if err := t.ExecuteTemplate(w, "top", dat); err != nil {
		fmt.Println(err)
	}

	//POSTメソッドのフォームをterminal上に表示
	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}
