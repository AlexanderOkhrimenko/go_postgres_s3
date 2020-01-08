package modules

import "fmt"

// Добавляет новую запись в таблицу  возвращает id
func InsertDBurl(url string) uint64 {

	var lastInsertId uint64
	err := db.QueryRow("INSERT INTO encore_tab (url) VALUES ('dfgdfg') returning id;",
		url).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsertId)
	return lastInsertId

}