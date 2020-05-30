package dbctl

import (
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

// DeleteBook は本の削除を行う関数です
func DeleteBook(apiID string) error {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Println(apiID)
	b := ""
	rows, err := db.Query("select book_info_id from book_info where api_id = ?", apiID)
	if err != nil {
		log.Println(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&b)
		if err != nil {
			log.Printf(errFormat, err, f.Name(), file, line)
			return err
		}

		break
	}

	fmt.Println(f.Name(), b)

	del, err := db.Query("delete from book_statuses where book_info_id = ?", b)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer del.Close()

	fmt.Println(rows)

	return nil

}

// // DeleteBook は本の削除を行う関数です
// func DeleteBook(apiID string) error {
// 	pc, file, line, _ := runtime.Caller(0)
// 	f := runtime.FuncForPC(pc)
// 	fmt.Println(apiID)
// 	b := ""
// 	rows, err := db.Query("select book_info_id from book_info where apiID = ?", apiID)
// 	if err != nil {
// 		log.Println(errFormat, err, f.Name(), file, line)
// 		return err
// 	}
// 	rows.Next()
// 	err = rows.Scan(&b)
// 	log.Println(b)
// 	return err
// }
