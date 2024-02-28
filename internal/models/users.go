package models

type User struct {
	ID             int
	Username       string
	Email          string
	DocumentNumber int64
	Phone          int64
	PasswordHash   string
}