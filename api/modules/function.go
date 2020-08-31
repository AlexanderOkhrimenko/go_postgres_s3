package modules

import "context"

// Adds a new record to the table and returns the id
func InsertDBurl(url string) uint64 {

	conn := ConnectPG()
	defer ConnectPG().Close(context.Background())

	var lastInsertId uint64
	err := conn.QueryRow(context.Background(), "INSERT INTO encore_tab (url) VALUES ($1) returning id;",
		url).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	return lastInsertId
}
