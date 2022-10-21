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
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	query := "CREATE TABLE IF NOT EXISTS sessions(sessionname VARCHAR(40) PRIMARY KEY, username VARCHAR(40));"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	query = "CREATE TABLE IF NOT EXISTS admins(adminname VARCHAR(40) PRIMARY KEY, password VARCHAR(40), status VARCHAR(40));"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM admins")
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
