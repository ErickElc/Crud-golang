package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDataBase() *sql.DB {
	connect := "user=postgres dbname=postgres password=admin host=localhost sslmode=disable port=8000"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}
