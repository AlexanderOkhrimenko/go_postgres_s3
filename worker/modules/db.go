package modules

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"os"
)

func ConnectPG() (conn *pgx.Conn) {

	pgDBName := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")

	fmt.Println(pgDBName, pgUser, pgPass)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgUser, pgPass, "postgresql", "5432", pgDBName, "disable")

	fmt.Println(connStr, "---------- DB")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You connect to your database")

	return conn
}
