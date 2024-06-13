package builders

import (
	"VieiraDJS/app/helpers/cryptography"
	"VieiraDJS/app/helpers/validators"
	"VieiraDJS/app/models"
)

func NewUser(username string, password string, email string) (*validators.ValidatedUser, error) {
	salt, _ := cryptography.GenerateSalt(16)
	hashedPassword, _ := cryptography.HashPassword(password, salt)

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Salt: salt,
		Email: email,
	}

	validated_user := &validators.ValidatedUser{
		User: *user,
	}

	if err := validated_user.ValidateUser(); err != nil {
		return nil, err
	}

	return validated_user, nil
}
