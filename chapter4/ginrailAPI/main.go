package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStation(ctx *gin.Context) {
	var station StationResource
	id := ctx.Param("station_id")
	err := DB.QueryRow("SELECT ID, NAME, CAST(OPENING_TIME AS CHAR), CAST(CLOSING_TIME AS CHAR) FROM station WHERE id = ?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"result": station,
	})
}

func CreateStation(ctx *gin.Context) {
	var station StationResource
	if err := ctx.BindJSON(&station); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	stmt, err := DB.Prepare(`
		INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME) VALUES (?,?,?)
	`)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := stmt.Exec(station.Name, station.OpeningTime, station.ClosingTime)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	newID, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	station.ID = int(newID)
	ctx.JSON(201, gin.H{
		"result": station,
	})
}

func RemoveStation(ctx *gin.Context) {
	id := ctx.Param("station_id")
	stmt, err := DB.Prepare("DELETE FROM station WHERE id = ?")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(200)
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "../railAPI/railapi.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r := gin.Default()
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8080")
}
