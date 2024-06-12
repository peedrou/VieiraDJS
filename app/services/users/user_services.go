package users

import (
	"VieiraDJS/app/helpers/cryptography"
	"VieiraDJS/app/models"
	"errors"

	"github.com/gocql/gocql"
)

func RegisterUser(session *gocql.Session, username, password, email string) error {
	salt, err := cryptography.GenerateSalt(16)
	if err != nil {
		return err
	}

	hashedPassword, err := cryptography.HashPassword(password, salt)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
		Salt:     salt,
		Email:    email,
	}

	if err := session.Query(`INSERT INTO users (username, password, salt, email) VALUES (?, ?, ?, ?)`,
		user.Username, user.Password, user.Salt, user.Email).Exec(); err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(session *gocql.Session, username string, password string) (bool, error) {
	var storedPassword, salt string
	if err := session.Query(`SELECT password, salt FROM users WHERE username = ?`, username).Scan(&storedPassword, &salt); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return false, nil
		}
		return false, err
	}

	if err := cryptography.ComparePassword(storedPassword, password, salt); err != nil {
		return false, nil
	}

	return true, nil
}
