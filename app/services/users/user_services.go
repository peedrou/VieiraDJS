package users

import (
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/cryptography"
	"errors"

	"github.com/gocql/gocql"
)

func RegisterUser(session *gocql.Session, username string, password string, email string) error {

	validated_user, errUserCreation := builders.NewUser(
		username, password, email,
	)

	if errUserCreation != nil {
		return errors.New(errUserCreation.Error())
	}

	// TODO: Create other validating steps (check if other user exists with the emauil, username, etc)
	if err := session.Query(
		`INSERT INTO users (username, password, salt, email) VALUES (?, ?, ?, ?)`,
		validated_user.User.Password,
		validated_user.User.Password,
		validated_user.User.Salt,
		validated_user.User.Email,
	).Exec(); err != nil {
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
