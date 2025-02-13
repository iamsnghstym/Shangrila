package main

import (
	"database/sql"
	"fmt"
	branca "github.com/hako/branca"
	"log"
)


const (
	databaseURL = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	port = 26257
)

func main() {
	fmt.Println("Welcome to Shangrila.")

	// Open database connection
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Could not open database connection :%v/n", err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database :%v/n", err)
	}
}