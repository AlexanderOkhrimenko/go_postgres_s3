package modules

// Добавляет новую запись в таблицу  возвращает id
func InsertDBurl(url string) uint64 {

	var lastInsertId uint64
	err := db.QueryRow("INSERT INTO encore_tab (url) VALUES ($1) returning id;",
		url).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	return lastInsertId
}
