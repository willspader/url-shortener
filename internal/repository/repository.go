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

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
