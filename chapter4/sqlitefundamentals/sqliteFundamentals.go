package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS books(
			id INTEGER PRIMARY KEY, 
			isbn INTEGER,
			author VARCHAR(64),
			name VARCHAR(64) NULL
		)
	`)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
	stmt.Exec()

	// create
	stmt, err = db.Prepare(`
		INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
	stmt.Exec("Practical Go", "Anonymus", 1231231)
	log.Println("Successfully inserted the book in database!")
	// Read
	rows, err := db.Query(`SELECT id, name, author FROM books`)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s", tempBook.id, tempBook.name, tempBook.author)
	}
	// update
	stmt, err = db.Prepare(`
		UPDATE books SET name = ? WHERE id = ?
	`)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
	stmt.Exec("Practical Go updated", 1)
	log.Println("Successfully updated the book in database!")
	// delete
	stmt, err = db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
	stmt.Exec(1)
	log.Println("Successfully deleted the book in database!")
}
