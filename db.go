package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Db struct {
	connection *sql.DB
	Path       string
}

// Connects to the sqlite database in "path"/
func NewDb(path string) (*Db, error) {
	db := Db{}
	var err error
	db.connection, err = sql.Open("sqlite3", path)
	db.Path = path
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *Db) test() {
	defer db.connection.Close()

	sqlStmt := `
	SELECT content FROM sport_summary ORDER BY start_time DESC;
	`

	rows, err := db.connection.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var content string
		err = rows.Scan(&content)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(content)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
