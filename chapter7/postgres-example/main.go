package main

import (
	"log"

	"github.com/SantiagoBedoya/restfull-webservices/chapter7/postgres-example/models"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	log.Println(db)
}
