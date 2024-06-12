package models

type User struct {
	Username string
	Password string
	Salt     string
	Email    string
}
