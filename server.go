package main

import (
	"encoding/json"
	"github.com/go-zoo/bone"
	"log"
	"net/http"
)

const username = ""
const database = ""
const hostname = ""
const port = 5432
const password = ""

func main() {
	mux := bone.New()
	mux.Get("/v1/cache/:key", http.HandlerFunc(getCache))
	mux.Post("/v1/cache", http.HandlerFunc(setCache))
	mux.Get("/v1/clear-cache/:key", http.HandlerFunc(clearCache))
	http.ListenAndServe(":8000", mux)
}

var dbh = DBConnect{Username: username, Database: database, Hostname: hostname, Port: port, Password: password}
var db = connectDB(dbh)

func getCache(rw http.ResponseWriter, req *http.Request) {
	key := bone.GetValue(req, "key")

	rw.Header().Set("Content-Type", "application/json")

	result, err := GetCache(db, key)

	if err != nil {
		log.Print(err)
	}

	r, err := json.Marshal(result)

	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("fail"))
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(r))
	}
}

func setCache(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var record Record
	var err error

	err = decoder.Decode(&record)

	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Print(err)
	}

	_, err = SetCache(db, record)

	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("fail"))
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	}
}

func clearCache(rw http.ResponseWriter, req *http.Request) {
	key := bone.GetValue(req, "key")

	_, err := ClearCache(db, key)

	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("fail"))
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	}
}
