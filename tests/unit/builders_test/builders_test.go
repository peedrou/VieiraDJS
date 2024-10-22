package builders_test

import (
	"VieiraDJS/app/helpers/builders"
	"testing"
	"time"
)

func TestJobBuilder(t *testing.T) {
	_, err := builders.NewJob(
		true,
		2,
		time.Now(),
		"PT3H",
	)

	if err != nil {
		t.Errorf("NewJob() returned an unexpected error: %v", err)
	}
}

func TestUserBuilder(t *testing.T) {
	validated_user, err := builders.NewUser(
		"test_user",
		"password",
		"testemail@email.com",
	)

	if err != nil {
		t.Errorf("NewUser() returned an unexpected error: %v", err)
	}

	if validated_user.User.Password == "password" {
		t.Errorf("password is not hashed")
	}
}