package service

import "github.com/jackc/pgx"

// isUniqueViolation checks for unique violation error from pgx
func isUniqueViolation(err error) bool {
	pgerr, ok := err.(pgx.PgError)
	return ok && pgerr.Code == "23505"
}