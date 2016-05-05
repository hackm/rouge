package db

import (
	"database/sql"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:!Qaz2wsx@/rouge?interpolateParams=true&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
