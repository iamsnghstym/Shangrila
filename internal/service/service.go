package service

import (
	"database/sql"
	"github.com/hako/branca"
)

// Service holds the actual business logic for the application
type Service struct {
	Db *sql.DB
	Branca *branca.Branca
}

// New initialises the service
func New(db *sql.DB, codec *branca.Branca) *Service {
	return &Service {
		Db: db,
		Branca: codec,
	}
}