package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SantiagoBedoya/restfull-webservices/chapter7/jsonstore/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type DBClient struct {
	db *gorm.DB
}

type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

func (driver *DBClient) GetUsersByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("first_name")
	var query = `SELECT * FROM user WHERE data->>'first_name'=?`
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(users)
	w.Write(respJSON)
}

func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)
	driver.db.First(&user, vars["id"])
	var userData interface{}
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbclient.GetUser)
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods("POST")
	r.HandleFunc("/v1/user", dbclient.GetUsersByFirstName).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
