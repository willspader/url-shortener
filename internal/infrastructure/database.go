package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnection() *sql.DB {

	db, err := sql.Open("mysql", "url_shortener_app:admin@localhost:3306")
	if err != nil {
		panic(err.Error()) // TODO better error handling
	}

	return db
}
