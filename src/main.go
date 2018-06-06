package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	dbuser = "postgres"
	dbpass = ""
	dbname = "APP_DB"
)

func main() {
	dbInit()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func dbInit() {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbuser, dbpass, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to db!")

}
