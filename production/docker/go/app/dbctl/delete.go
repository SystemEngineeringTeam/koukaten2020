package dbctl

import (
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

// DeleteBook は本の削除を行う関数です
func DeleteBook(id string) error {

	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	var b string

	fmt.Println(id)

	rows, err := db.Query("select book_info_id from book_info where api_id = ?", id)
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
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

	del, err := db.Prepare("delete from book_statuses where book_data_id = ?")
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer del.Close()

	result, err := del.Exec(b)

	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		fmt.Println(rowsAffect)
		return err
	}
	// rows, err := db.Query("delete from tasks where id = ?;", id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println(rows)

	return nil

}
