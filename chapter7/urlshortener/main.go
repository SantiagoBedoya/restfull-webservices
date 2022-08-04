package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SantiagoBedoya/restfull-webservices/chapter7/urlshortener/models"
	"github.com/SantiagoBedoya/restfull-webservices/chapter7/urlshortener/utils"
	"github.com/gorilla/mux"
)

type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	id := utils.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url from web_url WHERE id = $1", id).Scan(&url)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	responseMap := map[string]interface{}{"url": url}
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES ($1) RETURNING id", record.URL).Scan(&id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	responseMap := map[string]interface{}{"encoded_string": utils.ToBase62(id)}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
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
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
