package dbutils

import (
	"database/sql"
	"log"
	"os"
)

func Initialize(dbDriver *sql.DB) {
	stmt, err := dbDriver.Prepare(train)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	stmt, err = dbDriver.Prepare(station)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	stmt, err = dbDriver.Prepare(schedule)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("All tabled created/initialized successfully")
}
