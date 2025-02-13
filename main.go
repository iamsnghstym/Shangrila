package main

import (
	"database/sql"
	"fmt"
	"github.com/hako/branca"
	"github.com/iamsnghstym/Shangrila/internal/handler"
	"github.com/iamsnghstym/Shangrila/internal/service"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"net/http"
)


const (
	databaseURL = "postgresql://root@localhost:26257/shangrila?sslmode=disable"
	port = 3000
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

	// Initialise service
	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	svc := service.New(db, codec)

	// Initialise handler
	h := handler.New(svc)

	// server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Accepting connection on PORT -> %d\n", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}