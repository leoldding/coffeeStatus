package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	}

	//DELETE THIS
	query := "DROP TABLE IF EXISTS sessions"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	// create tables on initial setup
	query = "CREATE TABLE IF NOT EXISTS sessions(sessionname TEXT PRIMARY KEY, username VARCHAR(40), expiration TIMESTAMP WITH TIME ZONE);"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	query = "CREATE TABLE IF NOT EXISTS admins(adminname VARCHAR(40) PRIMARY KEY, password TEXT, status VARCHAR(40));"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	// add admin account on initial setup
	var count int
	query = "SELECT COUNT(*) FROM admins"
	row := db.QueryRow(query)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMINPASSWORD")), 8)
		query = fmt.Sprintf("INSERT INTO admins(adminname, password, status) VALUES ('%s', '%s', '%s');", os.Getenv("ADMINNAME"), hashedPassword, os.Getenv("STATUS"))
		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
