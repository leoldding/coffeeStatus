package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "CREATE TABLE IF NOT EXISTS sessions(sessionname VARCHAR(40) PRIMARY KEY, username VARCHAR(40));"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	query = "CREATE TABLE IF NOT EXISTS admins(adminname VARCHAR(40) PRIMARY KEY, password VARCHAR(40), status INTEGER);"
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
		query = "INSERT INTO admins(adminname, password) VALUES ('admin', 'password');"
		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
