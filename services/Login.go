package services

import (
	"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func VerifyLogin(usrname, pwd string) (token string, err error) {
	//fmt.Println(url, selection)
	db, err := sql.Open("sqlite3", "./big_ol.db")
	var username string
	var password string

	if err != nil {
		return "", err
	}
	q := `
	SELECT username, password FROM users WHERE username == ?
	`
	// select
	stmt, err := db.Prepare(q)
	defer stmt.Close()

	rows, err := stmt.Query(usrname)
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}

	for rows.Next() {
		rows.Scan(&username, &password)
		if pwd == password
	}
	return "", nil
}
