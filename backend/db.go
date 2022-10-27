package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var db *sql.DB
var err error

func initDB() {
	// compile all Postgres database information into string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	// connect to Postgres database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return
	}

	// create tables on initial setup
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sessions(sessionname TEXT PRIMARY KEY, username VARCHAR(40), expiration TIMESTAMP WITH TIME ZONE);")
	if err != nil {
		panic(err)
		return
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS admins(adminname VARCHAR(40) PRIMARY KEY, password TEXT);")
	if err != nil {
		panic(err)
		return
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS status(status VARCHAR(40), substatus TEXT);")
	if err != nil {
		panic(err)
		return
	}

	// add admin account on initial setup
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM admins;")
	err = row.Scan(&count)
	if err != nil {
		panic(err)
		return
	}
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMINPASSWORD")), 8)
		_, err = db.Exec("INSERT INTO admins(adminname, password) VALUES ($1, $2);", os.Getenv("ADMINNAME"), hashedPassword)
		if err != nil {
			panic(err)
			return
		}
	}

	// add initial status
	row = db.QueryRow("SELECT COUNT(*) FROM status;")
	err = row.Scan(&count)
	if err != nil {
		panic(err)
		return
	}
	if count == 0 {
		_, err = db.Exec("INSERT INTO status(status, substatus) VALUES ($1, $2);", "yes", "none")
		if err != nil {
			panic(err)
			return
		}
	}

	return
}
