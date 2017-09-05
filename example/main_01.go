package main

import (
	"fmt"
	"time"
)

const username = ""
const database = ""
const hostname = ""
const port = 5432
const password = ""

func main() {
	c := DBConnect{Username: username, Database: database, Hostname: hostname, Port: port, Password: password}
	dbh := connectDB(c)

	token := time.Now().UnixNano()

	var r Record
	r.Key = "Duc"
	r.Value = fmt.Sprintf(`{"id": 1, "gender": "male", "address": "%d Bedford Ave"}`, token)
	r.TTL = 30

	result, err := SetCache(dbh, r)
	//result, err := GetCache(dbh, "Duc")
	//result, err := ClearCache(dbh, "Duc")

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
