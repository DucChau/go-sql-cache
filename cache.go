package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DBConnect struct {
	Username string
	Database string
	Hostname string
	Port     int
	Password string
}

type Record struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	TTL       int    `json:"ttl"`
}

type Filters struct {
	Key string
}

func connectDB(c DBConnect) *sql.DB {
	connectionString := fmt.Sprintf("user=%v dbname=%v host=%v port=%v password=%v", c.Username, c.Database, c.Hostname, c.Port, c.Password)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Printf("Unable to connect to: %v, %v, %v, %v.", c.Hostname, c.Database, c.Port, c.Username)
	}

	return db
}

func GetCache(dbh *sql.DB, key string) (Record, error) {
	query := fmt.Sprintf(`SELECT id, created_at, key, value, ttl
				FROM sql_cache
				WHERE key = $1 AND extract(epoch from now()) < (extract(epoch from created_at) + ttl)
				ORDER by id desc limit 1`)

	var r Record

	err := dbh.QueryRow(query, key).Scan(&r.ID, &r.CreatedAt, &r.Key, &r.Value, &r.TTL)

	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Printf("Query - %s (%s).", query, key)
			return r, err
		}
	}

	return r, nil
}

func SetCache(dbh *sql.DB, r Record) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT into sql_cache (key, value, ttl)
				VALUES ($1, $2, $3)`)

	stmt, err := dbh.Prepare(query)

	if err != nil {
		log.Printf("Prepare error - %s", query)
		return nil, err
	}

	result, err := stmt.Exec(r.Key, r.Value, r.TTL)

	if err != nil {
		log.Printf("Query - %s (%s, %s, %d).", query, r.Key, r.Value, r.TTL)
		return result, err
	}

	return result, nil
}

func ClearCache(dbh *sql.DB, key string) (sql.Result, error) {
	query := fmt.Sprintf(`UPDATE sql_cache SET ttl = 0 WHERE key = $1`)

	stmt, err := dbh.Prepare(query)

	if err != nil {
		log.Printf("Prepare error - %s", query)
		return nil, err
	}

	result, err := stmt.Exec(key)

	if err != nil {
		log.Printf("Query - %s (%s).", query, key)
		return result, err
	}

	return result, nil
}
