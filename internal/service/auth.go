package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	TokenLifeSpan = time.Hour * 24 * 14 // 14 days

)
// LoginOutput response
type LoginOutput struct {
	Token string
	ExpiresAt time.Time
	AuthUser User
}

// Login handles login fucntionality for the user (insecure login without password)
func (s * Service) Login(ctx context.Context, email string) (LoginOutput, error) {
	var out LoginOutput

	// Validate email
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return out, ErrEmailInvalid
	}

	query := "SELECT id, username FROM users WHERE email = $1"
	err := s.Db.QueryRowContext(ctx, query, email).Scan(&out.AuthUser.ID, &out.AuthUser.Username);

	if errors.Is(err, sql.ErrNoRows) {
		return out, ErrUserNotFound
	}

	if err != nil {
		return out, fmt.Errorf("could not query user: %v/n", err)
	}

	// Set tokens
	idStr := strconv.FormatInt(out.AuthUser.ID, 10)
	out.Token, err = s.Branca.EncodeToString(idStr)
	if err != nil {
	    return out, fmt.Errorf("could not create token for the user: %v/n", err)
	}

	out.ExpiresAt = time.Now().Add(TokenLifeSpan)
	return out, nil
}