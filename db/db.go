package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connect := "user=postgres dbname=store_db password=0611 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connect)

	if err != nil {
		panic(err.Error())
	}

	return db
}
