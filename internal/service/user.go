package service

import (
	"errors"
	"regexp"
)

var (
	// ErrUserNotFound when the user isn't found in the db
	ErrUserNotFound = errors.New("User not found")

	// Regular expression for email validation
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")

	// ErrEmailInvalid when email is invalid
	ErrEmailInvalid = errors.New("email is in an invalid format")
)
// User models the actual user
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}