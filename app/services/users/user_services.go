package users

import (
	"VieiraDJS/app/helpers/cryptography"
	"VieiraDJS/app/models"

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
