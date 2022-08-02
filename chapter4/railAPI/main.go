package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SantiagoBedoya/restfull-webservices/chapter4/dbutils"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

func (t TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE id = ?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
		return
	}
	response.WriteEntity(t)
}

func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
	stmt, err := DB.Prepare(`
		INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) VALUES (?,?)
	`)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	result, err := stmt.Exec(b.DriverName, b.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	} else {
		newID, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
			response.AddHeader("Content-Type", "text/plain")
			response.WriteErrorString(http.StatusInternalServerError, err.Error())
			return
		}
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	}
}

func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	stmt, err := DB.Prepare("DELETE FROM train WHERE id = ?")
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	response.WriteHeader(http.StatusOK)
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "railapi.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on port 8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
