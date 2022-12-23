package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnection() *sql.DB {

	db, err := sql.Open("mysql", "url_shortener_app:url_shortener_pwd@/url_shortener_db")
	if err != nil {
		panic(err.Error()) // TODO better error handling
	}

	return db
}
