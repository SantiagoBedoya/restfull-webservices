package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

func (c App) GetTrain() revel.Result {
	var train TrainResource
	id := c.Params.Route.Get("train-id")
	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "Logan"
	train.OperatingStatus = true
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

func (c App) CreateTrain() revel.Result {
	var train TrainResource
	c.Params.BindJSON(&train)
	train.ID = 2
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

func (c App) RemoveTrain() revel.Result {
	id := c.Params.Route.Get("train-id")
	log.Println("deleted ", id)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON("")
}

func (c App) Index() revel.Result {
	return c.Render()
}
