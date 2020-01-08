package modules

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func init() {

	pgDbName := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")

	pgDbName = "t1"
	pgUser = "apiuser"
	pgPass = "Wmf84vjKWp@f9okko23($F"

	fmt.Println(pgDbName, pgUser, pgPass)

	connStr := "host=postgresql user=" + pgUser + " password=" + pgPass + " dbname=" + pgDbName + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connect to your database")


}
