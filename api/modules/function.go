package modules

// Adds a new record to the table and returns the id
func InsertDBurl(url string) uint64 {

	var lastInsertId uint64
	err := db.QueryRow("INSERT INTO encore_tab (url) VALUES ($1) returning id;",
		url).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	return lastInsertId
}
