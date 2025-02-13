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

	db := mustInitDB()
	defer db.Close()

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	codec.SetTTL(uint32(service.TokenLifeSpan.Seconds()))

	svc := service.New(db, codec)
	h := handler.New(svc)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server running on PORT %d", port)
	log.Fatal(http.ListenAndServe(addr, h))
}

// mustInitDB opens the database and ensures it's accessible
func mustInitDB() *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal("Could not open database connection:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Could not ping database:", err)
	}

	return db
}