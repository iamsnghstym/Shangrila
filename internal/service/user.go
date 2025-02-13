package main

// User models the actual user
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}