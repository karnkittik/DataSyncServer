package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@/mydb?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
