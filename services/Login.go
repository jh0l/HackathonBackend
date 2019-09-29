package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"errors"

	"time"

	"fmt"

	"github.com/astaxie/beego"
)

func NewHashToken(username string) (token string, err error) {
	secret := beego.AppConfig.String("server_secret")
	key := []byte(secret)
	user := []byte(username)
	mac := hmac.New(sha256.New, key)
	mac.Write(user)
	hashStr := hex.EncodeToString(mac.Sum(nil))
	if err != nil {
		return "", err
	}
	fmt.Println("secret: %s\n, key: %s\n, user: %s\n, mac: %s\n, hashStr: %s\n", secret, key, user, mac, hashStr)
	return string(hashStr + " : " + string(time.Now().Format("12:54:04"))), nil
}

//func VerifyToken()

func VerifyLogin(usrname, pwd string) (token string, err error) {
	//fmt.Println(url, selection)
	db, err := sql.Open("sqlite3", "./big_ol.db")

	if err != nil {
		return "", err
	}

	var username string
	var password string

	q := `
	SELECT username, password FROM users WHERE username == ?
	`
	// select
	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Println(err.Error())
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

func NewAccount(usrname, pwd string) (token string, err error) {
	db, err := sql.Open("sqlite3", "./big_ol.db")

	currentTime := time.Now()

	timeFormatted := currentTime.Format("2019-09-29T12:54:04")

	q := `
	INSERT INTO "users"("id","username","password","date_created") VALUES (NULL,?,?,?);
	`

	stmt, err := db.Prepare(q)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec(usrname, pwd, timeFormatted)

	if err != nil {
		return "", err
	}

	return NewHashToken(usrname)
}
