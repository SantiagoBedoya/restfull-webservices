package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/chapter7?sslmode=disable")
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS WEB_URL(
			ID SERIAL PRIMARY KEY,
			URL TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := stmt.Exec()
	log.Println(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
