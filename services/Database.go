package services

import (
	"database/sql"

	// i want it anyway
	_ "github.com/mattn/go-sqlite3"
)

<<<<<<< HEAD
var Db *sql.DB

func InitDatabase() (err error) {
=======
// Db - database connection to be used in all other modules
var Db *sql.DB

// InitDatabase -  Initialize the database
func InitDatabase() error {
>>>>>>> 4a4f3a3195eb6df8218aa21a558bdf491e7d82d5

	Db, err := sql.Open("sqlite3", "file:./big_ol.db?_fk")
	if err != nil {
		return err
	}

	stmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		date_created DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS charge_points (
		id INTEGER PRIMARY KEY,
		owner INTEGER NOT NULL,
		address TEXT NOT NULL DEFAULT "",
		description TEXT NOT NULL DEFAULT "",
		available_time TEXT NOT NULL DEFAULT "{}",
		price_tiers TEXT NOT NULL DEFAULT "{}",
		FOREIGN KEY(owner) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS available_sockets (
		id INTEGER PRIMARY KEY,
		charge_point INTEGER NOT NULL,
		socket_type TEXT NOT NULL,
		maximum_current FLOAT NOT NULL,
		number_available NOT NULL DEFAULT 0,
		number_remaining NOT NULL DEFAULT 0,
		FOREIGN KEY(charge_point) REFERENCES charge_points(id)
	);

	CREATE TABLE IF NOT EXISTS available_batteries (
		id INTEGER PRIMARY KEY,
		charge_point INTEGER NOT NULL,
		battery_type TEXT NOT NULL,
		number_remaining NOT NULL DEFAULT 0,
		FOREIGN KEY(charge_point) REFERENCES charge_points(id)
	);

	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY,
		charge_point INTEGER NOT NULL,
		description TEXT NOT NULL DEFAULT "",
		price_tiers TEXT NOT NULL DEFAULT "{}",
		date_ends DATETIME NOT NULL,
		FOREIGN KEY(charge_point) REFERENCES charge_points(id)
	)
	`
	_, err = Db.Exec(stmt)
	if err != nil {
		return err
	}

	seed := `
	INSERT INTO "users"("id","username","password","date_created") VALUES (NULL,'jh0l','1234567','2019-09-29T12:54:04');
	`

	_, err = Db.Exec(seed)

<<<<<<< HEAD
	if err != nil {
		return err
	}
=======
	_, err = Db.Exec(stmt)
	return err
>>>>>>> 4a4f3a3195eb6df8218aa21a558bdf491e7d82d5

	return nil
}
