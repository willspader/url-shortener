package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func (repository Repository) NewId() int64 {
	result, err := repository.db.Exec("INSERT INTO URL_SHORTENER (LONG_URL) VALUES (null)")
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}

func (repository Repository) UpdateRecord(id int64, url string) {
	repository.db.Exec("UPDATE URL_SHORTENER SET LONG_URL = ? WHERE ID = ?", url, id)
}

func (repository Repository) GetLongUrl(id int64) string {
	var row *sql.Row = repository.db.QueryRow("SELECT LONG_URL FROM URL_SHORTENER WHERE ID = ?", id)

	var longUrlPtr *string
	row.Scan(&longUrlPtr)

	return *longUrlPtr
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
