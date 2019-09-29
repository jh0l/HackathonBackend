package services

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (err error) {
	db, err := sql.Open("sqlite3", "./big_ol.db")

	if err != nil {
		errors.New("database broke")
	}

	q := `
	CREATE TABLE IF NOT EXISTS users(
		Id TEXT NOT NULL PRIMARY KEY,
		Username TEXT,
		Password TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err = db.Prepare(q)

	return err
}
