package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"

	"errors"

	"fmt"

	"github.com/astaxie/beego"
)

func NewHashToken(username string) (token string, err error) {
	secret := beego.AppConfig.String("server_secret")
	key := []byte(secret)
	user := []byte(username)
	mac := hmac.New(sha256.New, key)
	mac.Write(user)
	hashStr, err := fmt.Printf("%x", mac)
	if err != nil {
		return "", err
	}
	fmt.Println("secret: %s\n, key: %s\n, user: %s\n, mac: %s\n, hashStr: %s\n", secret, key, user, mac, hashStr)
	return string(hashStr), nil
}

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
	if err != nil {
		fmt.Println(err.Error(), "qwrqe5t")
	}
	defer stmt.Close()

	rows, err := stmt.Query(usrname)
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}

	for rows.Next() {
		rows.Scan(&username, &password)
	}
	if pwd == password {
		hash, err := NewHashToken(username)
		if err != nil {
			return "", err
		}
		return hash, nil
	}

	return "", errors.New("password does not match")
}
