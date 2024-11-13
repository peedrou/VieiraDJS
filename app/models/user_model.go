package models

import (
	"github.com/gocql/gocql"
)

type User struct {
	UserID   gocql.UUID
	Username string
	Password string
	Salt     string
	Email    string
}
