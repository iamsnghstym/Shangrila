package main

import (
	"database/sql"
	"github.com/hako/branca"
)

// Service holds the actual business logic for the application
type Service struct {
	Db *sql.DB
	Branca *branca.Branca
}