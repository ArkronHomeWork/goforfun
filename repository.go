package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func databaseConnect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS gouser (username varchar , user_password varchar )")
	if err != nil {
		return nil, err
	}
	log.Printf("table gouser created with result %+v", res)
	return db, err
}
