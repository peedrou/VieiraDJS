package validators

import (
	"VieiraDJS/app/models"
	"errors"
)

type ValidatedUser struct {
	User models.User
}

func (u ValidatedUser) ValidateUser() error {
	if u.User.Username == "" || u.User.Password == "" || u.User.Email == "" {
		return errors.New("missing mandatory fields")
	}

	return nil
}