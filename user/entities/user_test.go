package entities_test

import (
	"gomq/user/entities"
	"testing"
)

func TestNewUs(t *testing.T) {
	user, err := entities.NewUser("Paulo")
	if err != nil {
		t.Errorf("Error not add new user %v", err)
	}
	t.Logf("Data %v", user)
}
