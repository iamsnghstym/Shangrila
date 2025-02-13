package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// ErrUserNotFound when the user isn't found in the db
	ErrUserNotFound = errors.New("User not found")

	// Regular expression for email validation
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")

	// Regular expression for email validation
	rxUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{0,17}$")


	// ErrEmailInvalid when email is invalid
	ErrEmailInvalid = errors.New("email is in an invalid format")

	// ErrEmailAlreadyTaken when email is already present
	ErrEmailAlreadyTaken = errors.New("email is already taken")

	// ErrUsernameInvalid when email is invalid
	ErrUsernameInvalid = errors.New("username is invalid")

	// ErrUsernameAlreadyTaken when email is already present
	ErrUsernameAlreadyTaken = errors.New("username is already taken")
)

// User models the actual user
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}

// CreateUser inserts user in the db
func (s *Service) CreateUser(ctx context.Context, email, username string) error {
	// Validate email
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return ErrEmailInvalid
	}

	// Validate username
	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return ErrUsernameInvalid
	}


	query := "INSERT INTO users (email, username) VALUES ($1, $2)"
	_, err := s.Db.ExecContext(ctx, query, email, username)
	isUniqueError := isUniqueViolation(err)

	if isUniqueError && strings.Contains(err.Error(), "email") {
		return ErrEmailAlreadyTaken
	}

	if isUniqueError && strings.Contains(err.Error(), "username") {
		return ErrUsernameAlreadyTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}
	return nil
}
