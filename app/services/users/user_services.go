package users

import (
	crud "VieiraDJS/app/db/CRUD"
	"VieiraDJS/app/helpers/builders"
	"VieiraDJS/app/helpers/cryptography"
	"errors"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func RegisterUser(session *gocql.Session, username string, password string, email string) (gocql.UUID, error) {
	someUUID := uuid.New()
	gocqlUUID, _ := gocql.ParseUUID(someUUID.String())

	validated_user, errUserCreation := builders.NewUser(
		gocqlUUID, username, password, email,
	)

	if errUserCreation != nil {
		return gocqlUUID, errors.New(errUserCreation.Error())
	}

	// TODO: Create other validating steps (check if other user exists with the emauil, username, etc)
	fields := []string{}

	fields = append(fields, "user_id", "username", "password", "salt", "email")

	err := crud.CreateModel(
		session,
		"users",
		fields,
		gocqlUUID,
		validated_user.User.Username,
		validated_user.User.Password,
		validated_user.User.Salt,
		validated_user.User.Email)

	if err != nil {
		return gocqlUUID, fmt.Errorf("there was an error inserting the user in the database: %v", err)
	}

	return gocqlUUID, nil
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

func ReadUsersGetUserID(session *gocql.Session, keys []string, values ...interface{}) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no keys were provided for user reading")
	}

	result, err := crud.ReadModel(session, "users", []string{"user_id"}, keys, values...)

	if err != nil {
		return nil, fmt.Errorf("there was an error reading the user(s) from the database: %v", err)
	}

	return result, nil
}

func ReadUsersGetEmail(session *gocql.Session, keys []string, values ...interface{}) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no keys were provided for user reading")
	}

	result, err := crud.ReadModel(session, "users", []string{"email"}, keys, values...)

	if err != nil {
		return nil, fmt.Errorf("there was an error reading the user(s) from the database: %v", err)
	}

	return result, nil
}

func ReadUsersGetUsername(session *gocql.Session, keys []string, values ...interface{}) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no keys were provided for user reading")
	}

	result, err := crud.ReadModel(session, "users", []string{"username"}, keys, values...)

	if err != nil {
		return nil, fmt.Errorf("there was an error reading the user(s) from the database: %v", err)
	}

	return result, nil
}
